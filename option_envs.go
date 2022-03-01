package gex

var (
	_            = Envs
	_ option     = (*optionEnvs)(nil)
	_ pipeOption = (*optionEnvs)(nil)
)

type optionEnvs struct {
	envs []*env
}

// Envs 配置命令运行的环境变量
func Envs(envs ...*env) *optionEnvs {
	return &optionEnvs{
		envs: envs,
	}
}

func (e *optionEnvs) apply(options *options) (err error) {
	options.envs = e.envs

	return
}

func (e *optionEnvs) applyPipe(options *pipeOptions) {
	options.envs = e.envs
}
