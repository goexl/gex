package gex

var (
	_            = DisableSystemEnvs
	_ option     = (*optionSystemEnvs)(nil)
	_ pipeOption = (*optionSystemEnvs)(nil)
)

type optionSystemEnvs struct{}

// DisableSystemEnvs 关闭系统环境变量
func DisableSystemEnvs() *optionSystemEnvs {
	return &optionSystemEnvs{}
}

func (se *optionSystemEnvs) apply(options *options) {
	options.systemEnvs = false
}

func (se *optionSystemEnvs) applyPipe(options *pipeOptions) {
	options.systemEnvs = false
}
