package gex

import (
	`io`
)

var (
	_        = Quiet
	_ option = (*optionQuiet)(nil)
)

type optionQuiet struct{}

// Quiet 配置清除终端显示
func Quiet() *optionQuiet {
	return &optionQuiet{}
}

func (q *optionQuiet) apply(options *options) {
	options.outs = make([]io.Writer, 0)
	options.errs = make([]io.Writer, 0)
}
