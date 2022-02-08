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
		checker: &equalChecker{
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
		checker: &containsChecker{
			contains: contains,
			options:  _options,
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
		checker: &containsAllChecker{
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
		checker: &containsAnyChecker{
			items:   items,
			options: _options,
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
		checker: &regexpChecker{
			regexp:  regexp,
			options: _options,
		},
	}
}

func (c *optionChecker) apply(options *options) {
	options.async = true
	options.checker = c.checker
}
