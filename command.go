package gex

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"syscall"

	"github.com/goexl/exc"
	"github.com/goexl/gox/field"
	"github.com/goexl/guc"
)

type command struct {
	params *params

	cmd    *exec.Cmd
	waiter *guc.WaitGroup
	enter  byte
}

func newCommand(params *params) *command {
	return &command{
		params: params,
		enter:  '\n',
	}
}

func (c *command) Exec() (code int, err error) {
	// 当出错时，打印到控制台
	if !c.params.echo && c.params.pwe {
		output := ""
		c.params.collectors[pwe] = newTerminalCollector(&output)
		defer c.cleanup(&output, &code, &err)
	}

	// 通知
	defer c.notify(&code, err)

	// 将所有的收集器加入到通知器中
	for _, _collector := range c.params.collectors {
		if _notifier, ok := _collector.(notifier); ok {
			c.params.notifiers = append(c.params.notifiers, _notifier)
		}
	}

	// 真正执行命令
	code, err = c.exec()

	return
}

func (c *command) exec() (code int, err error) {
	// 创建命令
	c.make()
	// 处理输入/输出
	if err = c.io(); nil != err {
		return
	}
	// 等待命令结束并取得返回码
	code, err = c.run()

	return
}

func (c *command) make() {
	args := c.params.args.String()
	if nil == c.params.context {
		// nolint: gosec
		c.cmd = exec.Command(c.params.name, args...)
	} else {
		// nolint: gosec
		c.cmd = exec.CommandContext(c.params.context, c.params.name, args...)
	}

	// 配置运行时目录
	if "" != c.params.dir {
		c.cmd.Dir = c.params.dir
	}

	// 配置运行时的环境变量
	if c.params.system {
		c.cmd.Env = os.Environ()
	}
	c.cmd.Env = append(c.cmd.Env, c.params.envs...)
}

func (c *command) io() (err error) {
	// 设置输入流
	if nil != c.params.stdin {
		c.cmd.Stdin = c.params.stdin
	}

	// 读取输出流数据
	if sop, soe := c.cmd.StdoutPipe(); nil != soe {
		err = soe
	} else {
		go c.read(sop, stdout)
	}
	if nil != err {
		return
	}

	// 读取错误流数据
	if sep, see := c.cmd.StderrPipe(); nil != see {
		err = see
	} else {
		go c.read(sep, stderr)
	}
	if nil != err {
		return
	}

	if nil != c.params.checks {
		c.waiter = new(guc.WaitGroup)
		// 特别注意，检查器等待器，是检查两个：输出流和错误流，但是只需要其中一个检查器退出，所有检查器都不应该再继续执行
		c.waiter.Add(1)
	}

	return
}

func (c *command) run() (code int, err error) {
	if c.params.echo {
		fmt.Print(c.params.name)
		fmt.Print(space)
		fmt.Print(c.params.args.Cli())
		fmt.Println()
	}

	// 执行命令
	if err = c.cmd.Start(); nil != err {
		return
	}

	// 如果有检查器，等待检查器结束
	if nil != c.params.checks {
		c.waiter.Wait()
	}

	// 如果是同步模式，等待命令执行完成
	if !c.params.async {
		code, err = c.wait()
	}

	return
}

func (c *command) wait() (code int, err error) {
	// 取得退出代码
	if err = c.cmd.Wait(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			if status, statsOk := exitErr.Sys().(syscall.WaitStatus); statsOk {
				code = status.ExitStatus()
			}
		}
	}

	return
}

func (c *command) cleanup(output *string, code *int, err *error) {
	// 检查状态码
	if 0 != *code {
		*err = exc.NewException(*code, "程序异常退出", field.New("error", *output))
	}

	if nil != err && nil != *err && "" != *output {
		_, _ = os.Stderr.WriteString(*output)
	}
}

func (c *command) notify(code *int, err error) {
	for _, _notifier := range c.params.notifiers {
		_notifier.Notify(*code, err)
	}
}

func (c *command) read(pipe io.ReadCloser, stream string) {
	done := false
	reader := bufio.NewReader(pipe)
	line, err := reader.ReadString(c.enter)
	for nil == err {
		c.line(line, stream)
		if checked, _ := c.params.check(line); checked && !done {
			c.waiter.Done()
			done = true
		}
		line, err = reader.ReadString(c.enter)
	}

	if nil != c.waiter {
		c.waiter.Done()
	}
}

func (c *command) line(line string, stream string) {
	// 如果需要回显，不使用日志输出，使用最原始的打包语句
	if c.params.echo {
		fmt.Println(line)
	}

	// 收集器
	for _, _collector := range c.params.collectors {
		_ = _collector.Collect(line, stream)
	}

	// 计数器
	for _, _counter := range c.params.counters {
		_ = _counter.Count(line, stream)
	}
}
