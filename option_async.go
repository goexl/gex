package exec

var (
	_        = Async
	_        = Sync
	_ option = (*optionAsync)(nil)
)

type optionAsync struct {
	async bool
}

// Async 配置异步执行
func Async() *optionAsync {
	return &optionAsync{
		async: true,
	}
}

// Sync 配置同步执行
func Sync() *optionAsync {
	return &optionAsync{
		async: false,
	}
}

func (a *optionAsync) apply(options *options) (err error) {
	options.async = a.async

	return
}
