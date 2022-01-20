package gex

import (
	`bufio`
	`fmt`
	`io`
	`os`
	`os/exec`
	`syscall`

	`github.com/storezhang/guc`
)

const enterChar = '\n'

// Run 执行外部命令
func Run(command string, opts ...option) (code int, err error) {
	_options := defaultOptions()
	for _, opt := range opts {
		opt.apply(_options)
	}

	// 通知
	defer notify(_options, &code, err)

	// 当出错时，打印到控制台
	if _options.pwe {
		output := ``
		_options.collectors[keyPwe] = newOutputStringCollector(&output, _options.max)
		defer printWhenError(&output, err)
	}

	// 将所有的收集器加入到通知器中
	for _, _collector := range _options.collectors {
		_options.notifiers = append(_options.notifiers, _collector)
	}

	// 创建真正的命令
	cmd := exec.Command(command, _options.args...)
	// 配置运行时目录
	if `` != _options.dir {
		cmd.Dir = _options.dir
	}

	// 配置运行时的环境变量
	if _options.systemEnvs {
		cmd.Env = os.Environ()
	}
	for _, _env := range _options.envs {
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

	var checkerGroup *guc.WaitGroup
	if nil != _options.checker {
		checkerGroup = new(guc.WaitGroup)
		// 特别注意，检查器等待器，是检查两个：输出流和错误流，但是只需要其中一个检查器退出，所有检查器都不应该再继续执行
		checkerGroup.Add(1)
	}

	// 读取输出流数据
	go read(stdout, checkerGroup, CollectorModeStdout, _options)
	// 读取错误流数据
	go read(stderr, checkerGroup, CollectorModeStderr, _options)

	// 如果有检查器，等待检查器结束
	if nil != _options.checker {
		checkerGroup.Wait()
	}

	// 如果是同步模式，等待命令执行完成
	if !_options.async {
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

func printWhenError(output *string, err error) {
	if nil != err {
		_, _ = os.Stderr.WriteString(*output)
	}
}

func notify(options *options, code *int, err error) {
	for _, _notifier := range options.notifiers {
		_notifier.notify(*code, err)
	}
}

func read(pipe io.ReadCloser, checker *guc.WaitGroup, mode CollectorMode, options *options) {
	done := false
	reader := bufio.NewReader(pipe)
	line, err := reader.ReadString(enterChar)
	for nil == err {
		collect(line, mode, options)

		if nil != options.checker {
			if checked, _ := options.checker.check(line); checked && !done {
				checker.Done()
				done = true
			}
		}
		line, err = reader.ReadString(enterChar)
	}

	// 因为Checker只有一个，所以调用Done时必须先判断是不是整体已经结束，不然会导致计数为负
	if nil != checker && !checker.Completed() {
		checker.Done()
	}
}

func collect(line string, mode CollectorMode, options *options) {
	for _, _collector := range options.collectors {
		_ = _collector.collect(line, mode)
	}
}
