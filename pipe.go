package exec

import (
	`fmt`
	`os`
	`os/exec`
)

type pipe struct {
	name    string
	options *pipeOptions
}

func newPipe(name string, opts ...pipeOption) *pipe {
	_options := defaultPipeOptions()
	for _, opt := range opts {
		opt.applyPipe(_options)
	}

	return &pipe{
		name:    name,
		options: _options,
	}
}

func (p *pipe) cmd() (cmd *exec.Cmd) {
	if nil == p.options.context {
		cmd = exec.Command(p.name, p.options.args...)
	} else {
		cmd = exec.CommandContext(p.options.context, p.name, p.options.args...)
	}

	// 配置运行时目录
	if `` != p.options.dir {
		cmd.Dir = p.options.dir
	}

	// 配置运行时的环境变量
	if p.options.systemEnvs {
		cmd.Env = os.Environ()
	}
	for _, _env := range p.options.envs {
		cmd.Env = append(cmd.Env, fmt.Sprintf(`%s=%s`, _env.key, _env.value))
	}

	return
}
