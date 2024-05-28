package internal

import (
	"github.com/goexl/gex/internal/internal/constant"
	"github.com/goexl/gex/internal/param"
)

type Collect struct {
	params *param.Collect
}

func NewCollect(params *param.Collect) *Collect {
	return &Collect{
		params: params,
	}
}

func (c *Collect) Max(max int) (collect *Collect) {
	c.params.Max = max
	collect = c

	return
}

func (c *Collect) Stdout() (collect *Collect) {
	c.params.Stream = constant.Stdout
	collect = c

	return
}

func (c *Collect) Stderr() (collect *Collect) {
	c.params.Stream = constant.Stderr
	collect = c

	return
}

func (c *Collect) Trim() *Trim {
	return NewTrim(c)
}
