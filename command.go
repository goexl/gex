package gex

import (
	`bufio`
	`fmt`
	`io`
	`os`
	`os/exec`
	`sync`
	`syscall`
)

const enterChar = '\n'

var mutex sync.Mutex

type command struct {
	name    string
	options *options

	cmd     *exec.Cmd
	checker *waitGroup
}

func newCommand(name string, opts ...option) *command {
	mutex.Lock()
	defer mutex.Unlock()

	_options := defaultOptions()
	for _, opt := range opts {
		opt.apply(_options)
	}

	return &command{
		name:    name,
		options: _options,
	}
}

func (c *command) Exec() (code int, err error) {
	// 当出错时，打印到控制台
	if c.options.pwe {
		output := ``
		c.options.collectors[keyPwe] = newOutputStringCollector(&output, c.options.max)
		defer c.printWhenError(&output, &err)
	}

	// 通知
	defer c.notify(c.options, &code, err)

	// 将所有的收集器加入到通知器中
	for _, _collector := range c.options.collectors {
		if _notifier, ok := _collector.(notifier); ok {
			c.options.notifiers = append(c.options.notifiers, _notifier)
		}
	}

	// 真正执行命令
	code, err = c.exec()

	return
}

func (c *command) exec() (code int, err error) {
	// 创建命令
	c.make()
	// 配置运行时目录
	if `` != c.options.dir {
		c.cmd.Dir = c.options.dir
	}

	// 配置运行时的环境变量
	if c.options.systemEnvs {
		c.cmd.Env = os.Environ()
	}
	for _, _env := range c.options.envs {
		c.cmd.Env = append(c.cmd.Env, fmt.Sprintf(`%s=%s`, _env.key, _env.value))
	}

	// 设置输入流
	if nil != c.options.stdin {
		c.cmd.Stdin = c.options.stdin
	}

	// 找到输出流
	var stdout io.ReadCloser
	if stdout, err = c.cmd.StdoutPipe(); nil != err {
		return
	}

	// 找到错误流
	var stderr io.ReadCloser
	if stderr, err = c.cmd.StderrPipe(); nil != err {
		return
	}

	// 执行命令
	if err = c.cmd.Start(); nil != err {
		return
	}

	// 处理管理
	if err = c.pipe(); nil != err {
		return
	}

	if nil != c.options.checker {
		c.checker = new(waitGroup)
		// 特别注意，检查器等待器，是检查两个：输出流和错误流，但是只需要其中一个检查器退出，所有检查器都不应该再继续执行
		c.checker.Add(1)
	}

	// 读取输出流数据
	go c.read(stdout, OutputTypeStdout, c.options)
	// 读取错误流数据
	go c.read(stderr, OutputTypeStderr, c.options)

	// 如果有检查器，等待检查器结束
	if nil != c.options.checker {
		c.checker.Wait()
	}

	// 如果是同步模式，等待命令执行完成
	if !c.options.async {
		// 取得退出代码
		if err = c.cmd.Wait(); err != nil {
			if exitErr, ok := err.(*exec.ExitError); ok {
				if status, statsOk := exitErr.Sys().(syscall.WaitStatus); statsOk {
					code = status.ExitStatus()
				}
			}
		}
	}

	return
}

func (c *command) make() {
	if nil == c.options.context {
		c.cmd = exec.Command(c.name, c.options.args...)
	} else {
		c.cmd = exec.CommandContext(c.options.context, c.name, c.options.args...)
	}

	return
}

func (c *command) pipe() (err error) {
	pipes := c.options.pipes
	size := len(pipes)
	if 0 == size {
		return
	}

	commands := make([]*exec.Cmd, 0, size)
	// 创建管道的所有命令
	for index := 0; index < size; index++ {
		commands = append(commands, pipes[index].cmd())
	}

	// 依次将管理输入/输出连接起来
	for index := 0; index < size; index++ {
		if index == size-1 {
			c.cmd.Stdin, err = commands[index].StdoutPipe()
		} else {
			commands[index+1].Stdin, err = commands[index].StdoutPipe()
		}

		if nil != err {
			return
		}
	}

	// 第一个命令等待结束，其它命令异步执行
	for index := 0; index < size; index++ {
		cmd := commands[index]
		if 0 == index {
			err = cmd.Run()
		} else {
			err = cmd.Start()
		}

		if nil != err {
			return
		}
	}

	return
}

func (c *command) printWhenError(output *string, err *error) {
	if nil != err && nil != *err && `` != *output {
		_, _ = os.Stderr.WriteString(*output)
	}
}

func (c *command) notify(options *options, code *int, err error) {
	for _, _notifier := range options.notifiers {
		_notifier.Notify(*code, err)
	}
}

func (c *command) read(pipe io.ReadCloser, typ OutputType, options *options) {
	done := false
	reader := bufio.NewReader(pipe)
	line, err := reader.ReadString(enterChar)
	for nil == err {
		c.collect(line, typ, options)

		if nil != options.checker {
			if checked, _ := options.checker.Check(line); checked && !done {
				c.checker.Done()
				done = true
			}
		}
		line, err = reader.ReadString(enterChar)
	}

	if nil != c.checker {
		c.checker.Done()
	}
}

func (c *command) collect(line string, typ OutputType, options *options) {
	for _, _collector := range options.collectors {
		_ = _collector.Collect(line, typ)
	}
}
