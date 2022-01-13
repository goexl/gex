package gex

var _ option = (*optionChecker)(nil)

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
func ContainsChecker(contains string) *optionChecker {
	return &optionChecker{
		checker: &containsChecker{
			contains: contains,
		},
	}
}

// PathMatchChecker 路径匹配检查器
func PathMatchChecker(pattern string) *optionChecker {
	return &optionChecker{
		checker: &pathMatchChecker{
			pattern: pattern,
		},
	}
}

// RegexpChecker 正则表达式检查器
func RegexpChecker(regexp string) *optionChecker {
	return &optionChecker{
		checker: &regexpChecker{
			regexp: regexp,
		},
	}
}

func (c *optionChecker) apply(options *options) {
	options.checker = c.checker
}
