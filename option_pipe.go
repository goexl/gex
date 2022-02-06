package gex

var (
	_        = Pipe
	_ option = (*optionPipe)(nil)
)

type optionPipe struct {
	pipe []string
}

// Pipe 配置命令运行的参数
func Pipe(pipe ...string) *optionPipe {
	return &optionPipe{
		pipe: pipe,
	}
}

func (a *optionPipe) apply(options *options) {
	options.pipe = a.pipe
}
