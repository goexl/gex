package gex

import (
	"strings"
)

type checkerContainsAll struct {
	items []string
	params *checkerParams
	all     strings.Builder
}

func newContainsAllChecker(params *checkerParams, items ...string) *checkerContainsAll {
	return &checkerContainsAll{
		params: params,
		items: items,
	}
}

func (cac *checkerContainsAll) Check(line string) (checked bool, err error) {
	if cac.params.cache {
		cac.all.WriteString(line)
	}

	checked = cac.containsAll(line, cac.items...)
	if !checked && cac.params.cache {
		checked = cac.containsAll(cac.all.String(), cac.items...)
	}

	return
}

func (cac *checkerContainsAll) containsAll(content string, items ...string) (contains bool) {
	contains = true
	for _, item := range items {
		contains = contains && strings.Contains(content, item)
	}

	return
}
