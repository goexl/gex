package gex

type checkerBuilder struct {
	params *checkerParams
	builder *builder
}

func newCheckerBuilder(builder *builder) *checkerBuilder {
	return &checkerBuilder{
		params: newCheckerParams(),
		builder: builder,
	}
}

func (cb *checkerBuilder) Cached() *checkerBuilder{
	cb.params.cache=true

	return cb
}

func (cb *checkerBuilder) ContainsAll(items ...string) *builder{
	cb.builder.params.checker = newContainsAllChecker(cb.params, items...)

	return cb.builder
}

func (cb *checkerBuilder) ContainsAny(items ...string) *builder{
	cb.builder.params.checker = newContainsAnyChecker(cb.params, items...)

	return cb.builder
}

func (cb *checkerBuilder) Contains(contains string) *builder{
	cb.builder.params.checker = newContainsChecker(contains,cb.params)

	return cb.builder
}

func (cb *checkerBuilder) Equal(equal string) *builder{
	cb.builder.params.checker = newEqualChecker(equal,cb.params)

	return cb.builder
}

func (cb *checkerBuilder) Filepath(pattern string) *builder{
	cb.builder.params.checker = newFilepathChecker(pattern,cb.params)

	return cb.builder
}

func (cb *checkerBuilder) Regexp(regexp string) *builder{
	cb.builder.params.checker = newRegexpChecker(regexp,cb.params)

	return cb.builder
}
