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
		c.options.notifiers = append(c.options.notifiers, _collector)
	}

	// 创建真正的命令
	cmd := exec.Command(c.name, c.options.args...)
	// 配置运行时目录
	if `` != c.options.dir {
		cmd.Dir = c.options.dir
	}

	// 配置运行时的环境变量
	if c.options.systemEnvs {
		cmd.Env = os.Environ()
	}
	for _, _env := range c.options.envs {
		cmd.Env = append(cmd.Env, fmt.Sprintf(`%s=%s`, _env.key, _env.value))
	}

	// 找到输出流
	var stdout io.ReadCloser
	if stdout, err = cmd.StdoutPipe(); nil != err {
		return
	}

	// 找到错误流
	var stderr io.ReadCloser
	if stderr, err = cmd.StderrPipe(); nil != err {
		return
	}

	// 执行命令
	if err = cmd.Start(); nil != err {
		return
	}

	if nil != c.options.checker {
		c.checker = new(waitGroup)
		// 特别注意，检查器等待器，是检查两个：输出流和错误流，但是只需要其中一个检查器退出，所有检查器都不应该再继续执行
		c.checker.Add(1)
	}

	// 读取输出流数据
	go c.read(stdout, CollectorTypeStdout, c.options)
	// 读取错误流数据
	go c.read(stderr, CollectorTypeStderr, c.options)

	// 如果有检查器，等待检查器结束
	if nil != c.options.checker {
		c.checker.Wait()
	}

	// 如果是同步模式，等待命令执行完成
	if !c.options.async {
		// 取得退出代码
		if err = cmd.Wait(); err != nil {
			if exitErr, ok := err.(*exec.ExitError); ok {
				if status, statsOk := exitErr.Sys().(syscall.WaitStatus); statsOk {
					code = status.ExitStatus()
				}
			}
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
		_notifier.notify(*code, err)
	}
}

func (c *command) read(pipe io.ReadCloser, typ CollectorType, options *options) {
	done := false
	reader := bufio.NewReader(pipe)
	line, err := reader.ReadString(enterChar)
	for nil == err {
		c.collect(line, typ, options)

		if nil != options.checker {
			if checked, _ := options.checker.check(line); checked && !done {
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

func (c *command) collect(line string, typ CollectorType, options *options) {
	for _, _collector := range options.collectors {
		_ = _collector.collect(line, typ)
	}
}
