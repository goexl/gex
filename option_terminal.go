package gex

import (
	`os`
)

var (
	_        = Terminal
	_ option = (*optionTerminal)(nil)
)

type optionTerminal struct{}

// Terminal 配置终端显示
func Terminal() *optionTerminal {
	return &optionTerminal{}
}

func (t *optionTerminal) apply(options *options) {
	options.outs = append(options.outs, os.Stdout)
	options.errs = append(options.errs, os.Stderr)
}
