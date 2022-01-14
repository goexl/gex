package gex

var (
	_        = Checker
	_        = ContainsChecker
	_        = PathMatchChecker
	_        = RegexpChecker
	_ option = (*optionChecker)(nil)
)

type optionChecker struct {
	checker checker
}

// Checker 配置检查器
func Checker(checker checker) *optionChecker {
	return &optionChecker{
		checker: checker,
	}
}

// ContainsChecker 字符串包含检查器
func ContainsChecker(contains string, opts ...checkerOption) *optionChecker {
	_options := defaultCheckerOptions()
	for _, opt := range opts {
		opt.applyChecker(_options)
	}

	return &optionChecker{
		checker: &containsChecker{
			contains: contains,
			cache:    _options.cache,
		},
	}
}

// PathMatchChecker 路径匹配检查器
func PathMatchChecker(pattern string, opts ...checkerOption) *optionChecker {
	_options := defaultCheckerOptions()
	for _, opt := range opts {
		opt.applyChecker(_options)
	}

	return &optionChecker{
		checker: &pathMatchChecker{
			pattern: pattern,
			cache:   _options.cache,
		},
	}
}

// RegexpChecker 正则表达式检查器
func RegexpChecker(regexp string, opts ...checkerOption) *optionChecker {
	_options := defaultCheckerOptions()
	for _, opt := range opts {
		opt.applyChecker(_options)
	}

	return &optionChecker{
		checker: &regexpChecker{
			regexp: regexp,
			cache:  _options.cache,
		},
	}
}

func (c *optionChecker) apply(options *options) {
	options.checker = c.checker
}
