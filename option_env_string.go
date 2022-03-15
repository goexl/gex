package gex

var (
	_        = StringEnv
	_        = StringEnvs
	_ option = (*optionStringEnv)(nil)
)

type optionStringEnv struct {
	envs []string
}

// StringEnv 环境变量
func StringEnv(env string) *optionStringEnv {
	return &optionStringEnv{
		envs: []string{env},
	}
}

// StringEnvs 环境变量列表
func StringEnvs(envs ...string) *optionStringEnv {
	return &optionStringEnv{
		envs: envs,
	}
}

func (se *optionStringEnv) apply(options *options) (err error) {
	for _, _env := range se.envs {
		options.envs = append(options.envs, parseEnv(_env))
	}

	return
}

func (se *optionStringEnv) applyPipe(options *pipeOptions) {
	for _, _env := range se.envs {
		options.envs = append(options.envs, parseEnv(_env))
	}
}
