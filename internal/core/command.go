package core

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"syscall"

	"github.com/goexl/exception"
	"github.com/goexl/gex/internal/internal/collector"
	"github.com/goexl/gex/internal/internal/constant"
	"github.com/goexl/gex/internal/internal/core"
	"github.com/goexl/gex/internal/param"
	"github.com/goexl/gox/field"
	"github.com/goexl/guc"
)

type Command struct {
	params *param.Command

	cmd    *exec.Cmd
	waiter *guc.WaitGroup
	enter  byte
}

func NewCommand(params *param.Command) *Command {
	return &Command{
		params: params,
		enter:  '\n',
	}
}

func (c *Command) Exec() (handler core.Handler, err error) {
	// 当出错时，打印到控制台
	if !c.params.Echo && c.params.Pwe {
		output := ""
		c.params.Collectors[constant.Pwe] = collector.NewTerminal(&output)
		defer c.cleanup(&output, &handler, &err)
	}

	// 通知
	defer c.notify(&handler, err)

	// 将所有的收集器加入到通知器中
	for _, _collector := range c.params.Collectors {
		if _notifier, ok := _collector.(core.Notifier); ok {
			c.params.Notifiers = append(c.params.Notifiers, _notifier)
		}
	}

	// 真正执行命令
	handler, err = c.exec()

	return
}

func (c *Command) exec() (handler core.Handler, err error) {
	// 创建命令
	c.make()
	// 处理输入/输出
	if err = c.io(); nil == err {
		handler, err = c.run() // 等待命令结束并取得返回码
	}

	return
}

func (c *Command) make() {
	args := c.params.Arguments.String()
	if nil == c.params.Context {
		// nolint: gosec
		c.cmd = exec.Command(c.params.Name, args...)
	} else {
		// nolint: gosec
		c.cmd = exec.CommandContext(c.params.Context, c.params.Name, args...)
	}

	// 配置运行时目录
	if "" != c.params.Directory {
		c.cmd.Dir = c.params.Directory
	}

	// 配置运行时的环境变量
	if nil != c.params.Environment {
		if c.params.Environment.System {
			c.cmd.Env = os.Environ()
		}
		c.cmd.Env = append(c.cmd.Env, c.params.Environment.Environments...)
	}
}

func (c *Command) io() (err error) {
	// 设置输入流
	if nil != c.params.Stdin {
		c.cmd.Stdin = c.params.Stdin
	}

	// 读取输出流数据
	if sop, soe := c.cmd.StdoutPipe(); nil != soe {
		err = soe
	} else {
		go c.read(sop, constant.Stdout)
	}
	if nil != err {
		return
	}

	// 读取错误流数据
	if sep, see := c.cmd.StderrPipe(); nil != see {
		err = see
	} else {
		go c.read(sep, constant.Stderr)
	}
	if nil != err {
		return
	}

	if nil != c.params.Logics {
		c.waiter = new(guc.WaitGroup)
		// 特别注意，检查器等待器，是检查两个：输出流和错误流，但是只需要其中一个检查器退出，所有检查器都不应该再继续执行
		c.waiter.Add(1)
	}

	return
}

func (c *Command) run() (handler core.Handler, err error) {
	if c.params.Echo {
		fmt.Print(c.params.Name)
		fmt.Print(constant.Space)
		fmt.Print(c.params.Arguments.Cli())
		fmt.Println()
	}

	// 执行命令
	if err = c.cmd.Start(); nil != err {
		return
	}

	// 如果有检查器，等待检查器结束
	if nil != c.params.Logics {
		c.waiter.Wait()
	}

	// 如果是同步模式，等待命令执行完成
	if !c.params.Async {
		handler, err = c.wait()
	}

	return
}

func (c *Command) wait() (handler core.Handler, err error) {
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

func (c *Command) cleanup(output *string, code *int, err *error) {
	// 检查状态码
	if 0 != *code {
		*err = exception.New().Code(*code).Message("程序异常退出").Field(field.New("error", *output)).Build()
	}

	if nil != err && nil != *err && "" != *output {
		_, _ = os.Stderr.WriteString(*output)
	}
}

func (c *Command) notify(code *int, err error) {
	for _, notifier := range c.params.Notifiers {
		notifier.Notify(*code, err)
	}
}

func (c *Command) read(pipe io.ReadCloser, stream string) {
	done := false
	reader := bufio.NewReader(pipe)
	line, err := reader.ReadString(c.enter)
	for nil == err {
		c.line(line, stream)
		if checked, _ := c.params.Check(line); checked && !done {
			c.waiter.Done()
			done = true
		}
		line, err = reader.ReadString(c.enter)
	}

	if nil != c.waiter {
		c.waiter.Done()
	}
}

func (c *Command) line(line string, stream string) {
	// 如果需要回显，不使用日志输出，使用最原始的打包语句
	if c.params.Echo {
		fmt.Println(line)
	}

	// 收集器
	for _, clt := range c.params.Collectors {
		_ = clt.Collect(line, stream)
	}

	// 计数器
	for _, counter := range c.params.Counters {
		_ = counter.Count(line, stream)
	}
}
