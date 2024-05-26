package builder

import (
	"fmt"

	"github.com/goexl/gex/internal/core"
	"github.com/goexl/gex/internal/internal/collector"
	"github.com/goexl/gex/internal/internal/constant"
	"github.com/goexl/gex/internal/param"
)

type Collector struct {
	builder *Command
	params  *param.Collector
}

func NewCollector(builder *Command) (collector *Collector) {
	return &Collector{
		builder: builder,
		params:  param.NewCollector(),
	}
}

func (c *Collector) Max(max int) (collector *Collector) {
	c.params.Max = max
	collector = c

	return
}

func (c *Collector) Stdout() (collector *Collector) {
	c.params.Stream = constant.Stdout
	collector = c

	return
}

func (c *Collector) Stderr() (collector *Collector) {
	c.params.Stream = constant.Stderr
	collector = c

	return
}

func (c *Collector) Trim() *Trim {
	return NewTrim(c)
}

func (c *Collector) Strings(strings *[]string) *Command {
	return c.put(collector.NewStrings(strings, c.params))
}

func (c *Collector) String(strings *string) *Command {
	return c.put(collector.NewString(strings, c.params))
}

func (c *Collector) put(collector core.Collector) *Command {
	c.builder.params.Collectors[fmt.Sprintf(constant.Point, collector)] = collector

	return c.builder
}
