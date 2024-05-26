package checker

import (
	"strings"

	"github.com/goexl/gex/internal/param"
)

type Contains struct {
	contains string
	params   *param.Checker
	all      strings.Builder
}

func NewContains(contains string, params *param.Checker) *Contains {
	return &Contains{
		contains: contains,
		params:   params,
	}
}

func (c *Contains) Check(line string) (checked bool, err error) {
	if c.params.Cache {
		c.all.WriteString(line)
	}

	checked = strings.Contains(line, c.contains)
	if !checked && c.params.Cache {
		checked = strings.Contains(c.all.String(), c.contains)
	}

	return
}
