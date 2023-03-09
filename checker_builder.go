package gex

type checkerBuilder struct {
	params  *checkerParams
	builder *Builder
}

func newCheckerBuilder(builder *Builder) *checkerBuilder {
	return &checkerBuilder{
		params:  newCheckerParams(),
		builder: builder,
	}
}

func (cb *checkerBuilder) Cached() *checkerBuilder {
	cb.params.cache = true

	return cb
}

func (cb *checkerBuilder) Contains(contains string) *checkerBuilder {
	cb.builder.params.checker = newContainsChecker(contains, cb.params)

	return cb
}

func (cb *checkerBuilder) Equal(equal string) *checkerBuilder {
	cb.builder.params.checker = newEqualChecker(equal, cb.params)

	return cb
}

func (cb *checkerBuilder) Filepath(pattern string) *checkerBuilder {
	cb.builder.params.checker = newFilepathChecker(pattern, cb.params)

	return cb
}

func (cb *checkerBuilder) Regexp(regexp string) *checkerBuilder {
	cb.builder.params.checker = newRegexpChecker(regexp, cb.params)

	return cb
}

func (cb *checkerBuilder) Build() *Builder {
	return cb.builder
}
