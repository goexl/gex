package gex

import (
	"fmt"
)

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
	cb.params.stream = stdout

	return cb
}

func (cb *collectorBuilder) Stderr() *collectorBuilder {
	cb.params.stream = stderr

	return cb
}

func (cb *collectorBuilder) TrimSpace() *collectorBuilder {
	cb.params.trim.space = true

	return cb
}

func (cb *collectorBuilder) Trim(cut string) *collectorBuilder {
	cb.params.trim.all = cut

	return cb
}

func (cb *collectorBuilder) TrimLeft(cut string) *collectorBuilder {
	cb.params.trim.left = cut

	return cb
}

func (cb *collectorBuilder) TrimRight(cut string) *collectorBuilder {
	cb.params.trim.right = cut

	return cb
}

func (cb *collectorBuilder) Strings(strings *[]string) *Builder {
	return cb.put(newStringsCollector(strings, cb.params))
}

func (cb *collectorBuilder) String(strings *string) *Builder {
	return cb.put(newStringCollector(strings, cb.params))
}

func (cb *collectorBuilder) put(collector collector) *Builder {
	cb.builder.params.collectors[fmt.Sprintf(point, collector)] = collector

	return cb.builder
}
