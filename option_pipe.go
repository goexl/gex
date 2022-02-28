package exec

var (
	_        = Pipe
	_        = Cli
	_ option = (*optionPipe)(nil)
)

type optionPipe struct {
	pipe *pipe
}

// Pipe 管道
func Pipe(name string, opts ...pipeOption) *optionPipe {
	return &optionPipe{
		pipe: newPipe(name, opts...),
	}
}

func (p *optionPipe) apply(options *options) (err error) {
	options.pipes = append(options.pipes, p.pipe)

	return
}
