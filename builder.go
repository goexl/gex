package gex

import (
	"context"

	"github.com/goexl/gox"
)

type builder struct {
	params *params
}

func newBuilder() *builder {
	return &builder{
		params: newParams(),
	}
}

func (b *builder) Args(args *gox.Args) *builder {
	b.params.args = args

	return b
}

func (b *builder) Context(ctx context.Context) *builder {
	b.params.context = ctx

	return b
}

func (b *builder) Dir(dir string) *builder {
	b.params.dir = dir

	return b
}

func (b *builder) Env(key string, value string) *builder {
	b.params.envs = append(b.params.envs, gox.StringBuilder(key, equal, value).String())

	return b
}

func (b *builder) Async() *builder {
	b.params.async = true

	return b
}

func (b *builder) Charset(charset string) *builder {
	b.params.charset = charset

	return b
}

func (b *builder) Charset(charset string) *builder {
	b.params.charset = charset

	return b
}
