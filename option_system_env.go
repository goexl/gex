package gex

var (
	_            = DisableSystemEnv
	_ option     = (*optionSystemEnv)(nil)
	_ pipeOption = (*optionSystemEnv)(nil)
)

type optionSystemEnv struct{}

// DisableSystemEnv 关闭系统环境变量
func DisableSystemEnv() *optionSystemEnv {
	return &optionSystemEnv{}
}

func (se *optionSystemEnv) apply(options *options) (err error) {
	options.system.envs = false

	return
}

func (se *optionSystemEnv) applyPipe(options *pipeOptions) {
	options.system.envs = false
}
