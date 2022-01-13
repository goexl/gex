package gex

var _ option = (*optionDir)(nil)

type optionDir struct {
	dir string
}

// Dir 配置命令运行的目录
func Dir(dir string) *optionDir {
	return &optionDir{
		dir: dir,
	}
}

func (d *optionDir) apply(options *options) {
	options.dir = d.dir
}
