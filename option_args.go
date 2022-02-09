package gex

import (
	`fmt`
	`strings`
)

var (
	_        = Args
	_        = Cli
	_ option = (*optionArgs)(nil)
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
func Args(args ...interface{}) (a *optionArgs) {
	a = new(optionArgs)
	for _, arg := range args {
		switch arg.(type) {
		case int8, *int8, uint8, *uint8, int, *int, uint, *uint, int32, *int32, uint32, *uint32, int64, *int64, uint64, *uint64:
			a.args = append(a.args, fmt.Sprintf(`%d`, arg))
		case float32, *float32, float64, *float64:
			a.args = append(a.args, fmt.Sprintf(`%f`, arg))
		case bool, *bool:
			a.args = append(a.args, fmt.Sprintf(`%b`, arg))
		case string, *string:
			a.args = append(a.args, fmt.Sprintf(`%s`, arg))
		}
	}

	return
}

func (a *optionArgs) apply(options *options) {
	options.args = a.args
}
