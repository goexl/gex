package builder

import (
	"github.com/goexl/gex/internal/core"
	"github.com/goexl/gex/internal/internal/checker"
	"github.com/goexl/gex/internal/internal/collector"
	core2 "github.com/goexl/gex/internal/internal/core"
	"github.com/goexl/gex/internal/param"
	"github.com/goexl/gox"
)

type Checker struct {
	gox.CannotCopy

	operator core.Operator
	params   *param.Checker
	builder  *Command
}

func NewChecker(operator core.Operator, builder *Command) *Checker {
	return &Checker{
		operator: operator,
		params:   param.NewChecker(),
		builder:  builder,
	}
}

func (c *Checker) Cached() *Checker {
	c.params.Cache = true

	return c
}

func (c *Checker) Contains(contains string) *Command {
	check := core2.NewCheck(c.operator, checker.NewContains(contains, c.params))
	c.builder.params.Checks = append(c.builder.params.Checks, check)

	return c.builder
}

func (c *Checker) Equal(equal string) *Command {
	check := core2.NewCheck(c.operator, checker.NewEqual(equal, c.params))
	c.builder.params.Checks = append(c.builder.params.Checks, check)

	return c.builder
}

func (c *Checker) Filepath(pattern string) *Command {
	check := core2.NewCheck(c.operator, checker.NewFilepath(pattern, c.params))
	c.builder.params.Checks = append(c.builder.params.Checks, check)

	return c.builder
}

func (c *Checker) Regexp(regexp string) *Command {
	check := core2.NewCheck(c.operator, collector.NewRegexp(regexp, c.params))
	c.builder.params.Checks = append(c.builder.params.Checks, check)

	return c.builder
}
