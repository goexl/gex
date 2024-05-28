package checker

import (
	"strings"

	"github.com/goexl/gex/internal/param"
)

type Equal struct {
	equal string

	params *param.Check
	all    strings.Builder
}

func NewEqual(equal string, params *param.Check) *Equal {
	return &Equal{
		equal:  equal,
		params: params,
	}
}

func (e *Equal) Check(line string) (checked bool, err error) {
	if e.params.Cache {
		e.all.WriteString(line)
	}

	checked = e.equal == line
	if !checked && e.params.Cache {
		checked = e.all.String() == e.equal
	}

	return
}
