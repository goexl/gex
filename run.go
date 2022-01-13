package gex

import (
	`fmt`
	`io`
	`os`
	`os/exec`
)

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

	return
}
