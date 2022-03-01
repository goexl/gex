package gex

var (
	_               = Cache
	_ checkerOption = (*optionCache)(nil)
)

type optionCache struct{}

// Cache 开启缓存
func Cache() *optionCache {
	return &optionCache{}
}

func (c *optionCache) applyChecker(options *checkerOptions) {
	options.cache = true
}
