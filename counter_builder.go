package gex

type counterBuilder struct {
	builder *Builder
	params  *counterParams
}

func newCounterBuilder(builder *Builder) *counterBuilder {
	return &counterBuilder{
		builder: builder,
		params:  newCounterParams(),
	}
}

func (cb *counterBuilder) Stdout() *counterBuilder {
	cb.params.stream = keyStdout

	return cb
}

func (cb *counterBuilder) Stderr() *counterBuilder {
	cb.params.stream = keyStderr

	return cb
}

func (cb *counterBuilder) Lines(total *int) *Builder {
	cb.builder.params.counters = append(cb.builder.params.counters, newLineCounter(total, cb.params))

	return cb.builder
}
