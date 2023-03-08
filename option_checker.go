package gex

var (
	_        = Checker
	_        = EqualChecker
	_        = ContainsChecker
	_        = ContainsAllChecker
	_        = ContainsAnyChecker
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

// EqualChecker 字符串全等检查器
func EqualChecker(equal string, opts ...checkerOption) *optionChecker {
	_options := defaultCheckerOptions()
	for _, opt := range opts {
		opt.applyChecker(_options)
	}

	return &optionChecker{
		checker: &checkerEqual{
			equal:   equal,
			options: _options,
		},
	}
}

// ContainsChecker 字符串包含检查器
func ContainsChecker(contains string, opts ...checkerOption) *optionChecker {
	_options := defaultCheckerOptions()
	for _, opt := range opts {
		opt.applyChecker(_options)
	}

	return &optionChecker{
		checker: &checkerContains{
			contains: contains,
			params:   _options,
		},
	}
}

// ContainsAllChecker 字符串包含检查器
func ContainsAllChecker(items []string, opts ...checkerOption) *optionChecker {
	_options := defaultCheckerOptions()
	for _, opt := range opts {
		opt.applyChecker(_options)
	}

	return &optionChecker{
		checker: &checkerContainsAll{
			items:   items,
			options: _options,
		},
	}
}

// ContainsAnyChecker 字符串包含检查器
func ContainsAnyChecker(items []string, opts ...checkerOption) *optionChecker {
	_options := defaultCheckerOptions()
	for _, opt := range opts {
		opt.applyChecker(_options)
	}

	return &optionChecker{
		checker: &checkerContainsAny{
			items:  items,
			params: _options,
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
		checker: &checkerFilepathMatch{
			pattern: pattern,
			options: _options,
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
		checker: &checkerRegexp{
			regexp:  regexp,
			options: _options,
		},
	}
}

func (c *optionChecker) apply(options *options) (err error) {
	options.async = true
	options.checker = c.checker

	return
}
