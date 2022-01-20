package gex

import (
	`context`
)

var (
	_        = Context
	_ option = (*optionContext)(nil)
)

type optionContext struct {
	context context.Context
}

// Context 配置执行上下文
func Context(context context.Context) *optionContext {
	return &optionContext{
		context: context,
	}
}

func (c *optionContext) apply(options *options) {
	options.context = c.context
}
