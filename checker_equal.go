package gex

import (
	"strings"
)

type checkerEqual struct {
	equal string

	params *checkerParams
	all     strings.Builder
}

func newEqualChecker(equal string, params *checkerParams) *checkerEqual{
	return &checkerEqual{
		equal: equal,
		params: params,
	}
}

func (ec *checkerEqual) Check(line string) (checked bool, err error) {
	if ec.params.cache {
		ec.all.WriteString(line)
	}

	checked = ec.equal == line
	if !checked && ec.params.cache {
		checked = ec.all.String() == ec.equal
	}

	return
}
