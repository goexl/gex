package gex

import (
	"strings"
)

type checkerContains struct {
	contains string
	params   *checkerParams
	all      strings.Builder
}

func newContainsChecker(contains string, params *checkerParams) *checkerContains {
	return &checkerContains{
		contains: contains,
		params:   params,
	}
}

func (cc *checkerContains) Check(line string) (checked bool, err error) {
	if cc.params.cache {
		cc.all.WriteString(line)
	}

	checked = strings.Contains(line, cc.contains)
	if !checked && cc.params.cache {
		checked = strings.Contains(cc.all.String(), cc.contains)
	}

	return
}
