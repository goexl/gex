package gex

import (
	"github.com/goexl/gox"
)

type environmentBuilder struct {
	builder *Builder
	params  *environmentParams
}

func newEnvironmentBuilder(builder *Builder) *environmentBuilder {
	return &environmentBuilder{
		builder: builder,
		params:  newEnvironmentParams(),
	}
}

func (eb *environmentBuilder) Only() *environmentBuilder {
	eb.params.system = false

	return eb
}

func (eb *environmentBuilder) Kv(key string, value any) *environmentBuilder {
	eb.params.environments = append(eb.params.environments, gox.StringBuilder(key, equal, value).String())

	return eb
}

func (eb *environmentBuilder) String(environments ...string) *environmentBuilder {
	for _, environment := range environments {
		eb.params.environments = append(eb.params.environments, environment)
	}

	return eb
}

func (eb *environmentBuilder) Build() *Builder {
	eb.builder.params.environment = eb.params

	return eb.builder
}
