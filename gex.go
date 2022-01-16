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

// Run 执行外部命令
func Run(command string, opts ...option) (code int, err error) {
	_options := defaultOptions()
	for _, opt := range opts {
		opt.apply(_options)
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

	checkerGroup := new(sync.WaitGroup)
	// 特别注意，检查器等待器，是检查两个：输出流和错误流，但是只需要其中一个检查器退出，所有检查器都不应该再继续执行
	checkerGroup.Add(1)

	// 读取输出流数据
	go read(stdout, checkerGroup, readTypeStdout, _options)
	// 读取错误流数据
	go read(stderr, checkerGroup, readTypeStderr, _options)

	exitGroup := new(sync.WaitGroup)
	exitGroup.Add(2)

	// 如果有检查器，等待检查器结束
	go waitChecker(exitGroup, checkerGroup, _options)
	// 如果是同步模式，等待命令执行完成
	go waitCommand(exitGroup, cmd, &code, err, _options)
	exitGroup.Wait()

	return
}

func waitChecker(exit *sync.WaitGroup, checker *sync.WaitGroup, options *options) {
	defer func() {
		exit.Done()
	}()

	if nil == options.checker {
		return
	}
	checker.Wait()
}

func waitCommand(exit *sync.WaitGroup, cmd *exec.Cmd, code *int, err error, options *options) {
	defer func() {
		exit.Done()
	}()

	if options.async {
		return
	}

	// 取得退出代码
	if err = cmd.Wait(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			if status, statsOk := exitErr.Sys().(syscall.WaitStatus); statsOk {
				*code = status.ExitStatus()
				err = nil
			}
		}
	}
}

func read(pipe io.ReadCloser, checker *sync.WaitGroup, readType readType, options *options) {
	done := false
	reader := bufio.NewReader(pipe)
	line, err := reader.ReadString(enterChar)
	for nil == err {
		go write(line, readType, options)

		if nil != options.checker {
			if checked, _ := options.checker.check(line); checked && !done {
				checker.Done()
				done = true
			}
		}
		line, err = reader.ReadString(enterChar)
	}
}

func write(line string, readType readType, options *options) {
	var writers []io.Writer
	switch readType {
	case readTypeStderr:
		writers = options.errs
	default:
		writers = options.outs
	}

	for _, writer := range writers {
		_, _ = writer.Write([]byte(line))
	}
}
