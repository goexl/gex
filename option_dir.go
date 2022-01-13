package gex

var _ option = (*optionDir)(nil)

type optionDir struct {
	dir string
}

// Dir 配置命令运行的参数
func Dir(dir string) *optionDir {
	return &optionDir{
		dir: dir,
	}
}

func (a *optionDir) apply(options *options) {
	options.dir = a.dir
}
