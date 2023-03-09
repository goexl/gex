package gex

import (
	"context"

	"github.com/goexl/gox"
	"github.com/goexl/gox/args"
)

// Builder 构建器
type Builder struct {
	params *params
	_      gox.CannotCopy
}

func newBuilder(name string) *Builder {
	return &Builder{
		params: newParams(name),
	}
}

func (b *Builder) Args(args *args.Args) *Builder {
	b.params.args = args

	return b
}

func (b *Builder) Context(ctx context.Context) *Builder {
	b.params.context = ctx

	return b
}

func (b *Builder) Dir(dir string) *Builder {
	b.params.dir = dir

	return b
}

func (b *Builder) Env(key string, value string) *Builder {
	b.params.envs = append(b.params.envs, gox.StringBuilder(key, equal, value).String())

	return b
}

func (b *Builder) System() *Builder {
	b.params.system = true

	return b
}

func (b *Builder) Async() *Builder {
	b.params.async = true

	return b
}

func (b *Builder) Pwe() *Builder {
	b.params.pwe = true

	return b
}

func (b *Builder) Charset(charset string) *Builder {
	b.params.charset = charset

	return b
}

func (b *Builder) Checker() *checkerBuilder {
	return newCheckerBuilder(b)
}

func (b *Builder) Check(checker checker) *Builder {
	// b.params.checker=

	return b
}

func (b *Builder) Collector() *collectorBuilder {
	return newCollectorBuilder(b)
}

func (b *Builder) Collect(collector collector) *Builder {
	b.params.collectors[collector.Name()] = collector

	return b
}

func (b *Builder) Counter() *counterBuilder {
	return newCounterBuilder(b)
}

func (b *Builder) Count(counter counter) *Builder {
	b.params.counters = append(b.params.counters, counter)

	return b
}

func (b *Builder) Notify(notifier notifier) *Builder {
	b.params.notifiers = append(b.params.notifiers, notifier)

	return b
}
