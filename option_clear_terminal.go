package gex

import (
	`io`
)

var (
	_        = ClearTerminal
	_ option = (*optionClearTerminal)(nil)
)

type optionClearTerminal struct{}

// ClearTerminal 配置清除终端显示
func ClearTerminal() *optionClearTerminal {
	return &optionClearTerminal{}
}

func (t *optionClearTerminal) apply(options *options) {
	options.outs = make([]io.Writer, 0)
	options.errs = make([]io.Writer, 0)
}
