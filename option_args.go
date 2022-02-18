package gex

import (
	`strings`
)

var (
	_            = Args
	_            = Cli
	_ option     = (*optionArgs)(nil)
	_ pipeOption = (*optionArgs)(nil)
)

type optionArgs struct {
	args []string
}

// Cli 使用命令行参数
func Cli(cli string) *optionArgs {
	return &optionArgs{
		args: strings.Split(cli, ` `),
	}
}

// Args 配置命令运行的参数
func Args(args ...interface{}) *optionArgs {
	return &optionArgs{
		args: str(args...),
	}
}

func (a *optionArgs) apply(options *options) (err error) {
	options.args = a.args

	return
}

func (a *optionArgs) applyPipe(options *pipeOptions) {
	options.args = a.args
}
