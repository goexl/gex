package gex

type checkerBuilder struct {
	operator operator
	params   *checkerParams
	builder  *Builder
}

func newCheckerBuilder(operator operator, builder *Builder) *checkerBuilder {
	return &checkerBuilder{
		operator: operator,
		params:   newCheckerParams(),
		builder:  builder,
	}
}

func (cb *checkerBuilder) Cached() *checkerBuilder {
	cb.params.cache = true

	return cb
}

func (cb *checkerBuilder) Contains(contains string) *Builder {
	check := newCheck(cb.operator, newContainsChecker(contains, cb.params))
	cb.builder.params.checks = append(cb.builder.params.checks, check)

	return cb.builder
}

func (cb *checkerBuilder) Equal(equal string) *Builder {
	check := newCheck(cb.operator, newEqualChecker(equal, cb.params))
	cb.builder.params.checks = append(cb.builder.params.checks, check)

	return cb.builder
}

func (cb *checkerBuilder) Filepath(pattern string) *Builder {
	check := newCheck(cb.operator, newFilepathChecker(pattern, cb.params))
	cb.builder.params.checks = append(cb.builder.params.checks, check)

	return cb.builder
}

func (cb *checkerBuilder) Regexp(regexp string) *Builder {
	check := newCheck(cb.operator, newRegexpChecker(regexp, cb.params))
	cb.builder.params.checks = append(cb.builder.params.checks, check)

	return cb.builder
}
