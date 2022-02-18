package gex

var (
	_            = Dir
	_ option     = (*optionDir)(nil)
	_ pipeOption = (*optionDir)(nil)
)

type optionDir struct {
	dir string
}

// Dir 配置命令运行的目录
func Dir(dir string) *optionDir {
	return &optionDir{
		dir: dir,
	}
}

func (d *optionDir) apply(options *options) (err error) {
	options.dir = d.dir

	return
}

func (d *optionDir) applyPipe(options *pipeOptions) {
	options.dir = d.dir
}
