package gex

var (
	_        = CommandStdin
	_ option = (*optionStdin)(nil)
)

type optionStdin struct {
	typ  stdinType
	args []interface{}
}

// CommandStdin 配置命令管道
func CommandStdin(opts ...pipeOption) *optionStdin {
	return &optionStdin{
		typ: stdinTypeCommand,
		args: []interface{}{
			opts,
		},
	}
}

func (s *optionStdin) apply(options *options) {
	options.stdinType = s.typ
	options.stdinArgs = s.args
}
