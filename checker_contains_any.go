package gex

import (
	"strings"
)

type checkerContainsAny struct {
	items []string

	params *checkerParams
	all    strings.Builder
}

func newContainsAnyChecker(params *checkerParams, items ...string) *checkerContainsAny {
	return &checkerContainsAny{
		params: params,
		items:  items,
	}
}

func (cac *checkerContainsAny) Check(line string) (checked bool, err error) {
	if cac.params.cache {
		cac.all.WriteString(line)
	}

	checked = cac.containsAny(line, cac.items...)
	if !checked && cac.params.cache {
		checked = cac.containsAny(cac.all.String(), cac.items...)
	}

	return
}

func (cac *checkerContainsAny) containsAny(content string, items ...string) (contains bool) {
	for _, item := range items {
		contains = strings.Contains(content, item)
		if contains {
			break
		}
	}

	return
}
