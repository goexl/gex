package gex

var _ option = (*optionEnvs)(nil)

type optionEnvs struct {
	envs []*env
}

// Envs 配置命令运行的环境变量
func Envs(envs ...*env) *optionEnvs {
	return &optionEnvs{
		envs: envs,
	}
}

func (e *optionEnvs) apply(options *options) {
	options.envs = e.envs
}
