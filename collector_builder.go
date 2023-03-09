package gex

type collectorBuilder struct {
	builder *Builder
	params  *collectorParams
}

func newCollectorBuilder(builder *Builder) *collectorBuilder {
	return &collectorBuilder{
		builder: builder,
		params:  newCollectorParams(),
	}
}

func (cb *collectorBuilder) Max(max int) *collectorBuilder {
	cb.params.max = max

	return cb
}

func (cb *collectorBuilder) Stdout() *collectorBuilder {
	cb.params.stream = keyStdout

	return cb
}

func (cb *collectorBuilder) Stderr() *collectorBuilder {
	cb.params.stream = keyStderr

	return cb
}

func (cb *collectorBuilder) Trim() *collectorBuilder {
	cb.params.trim.all = true

	return cb
}

func (cb *collectorBuilder) TrimLeft() *collectorBuilder {
	cb.params.trim.left = true

	return cb
}

func (cb *collectorBuilder) TrimRight() *collectorBuilder {
	cb.params.trim.right = true

	return cb
}

func (cb *collectorBuilder) Strings(strings *[]string) *Builder {
	collector := newStringsCollector(strings, cb.params)
	cb.builder.params.collectors[collector.Name()] = collector

	return cb.builder
}

func (cb *collectorBuilder) String(strings *string) *Builder {
	collector := newStringCollector(strings, cb.params)
	cb.builder.params.collectors[collector.Name()] = collector

	return cb.builder
}
