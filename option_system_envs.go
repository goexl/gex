package gex

var (
	_        = DisableSystemEnvs
	_ option = (*optionSystemEnvs)(nil)
)

type optionSystemEnvs struct{}

// DisableSystemEnvs 关闭系统环境变量
func DisableSystemEnvs() *optionSystemEnvs {
	return &optionSystemEnvs{}
}

func (a *optionSystemEnvs) apply(options *options) {
	options.systemEnvs = false
}
