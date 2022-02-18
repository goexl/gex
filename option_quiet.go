package gex

var (
	_        = Quiet
	_ option = (*optionQuiet)(nil)
)

type optionQuiet struct{}

// Quiet 配置清除终端显示
func Quiet() *optionQuiet {
	return &optionQuiet{}
}

func (q *optionQuiet) apply(options *options) (err error) {
	delete(options.collectors, keyTerminal)

	return
}
