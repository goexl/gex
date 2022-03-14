package gex

var (
	_        = Env
	_ option = (*optionEnv)(nil)
)

type optionEnv struct {
	key   string
	value string
}

// Env 环境变量
func Env(key string, value string) *optionEnv {
	return &optionEnv{
		key:   key,
		value: value,
	}
}

func (e *optionEnv) apply(options *options) (err error) {
	options.envs = append(options.envs, &env{
		key:   e.key,
		value: e.value,
	})

	return
}

func (e *optionEnv) applyPipe(options *pipeOptions) {
	options.envs = append(options.envs, &env{
		key:   e.key,
		value: e.value,
	})
}
