package gex

import (
	`strings`
)

type containsAllChecker struct {
	items []string

	options *checkerOptions
	all     strings.Builder
}

func (cac *containsAllChecker) Check(line string) (checked bool, err error) {
	if cac.options.cache {
		cac.all.WriteString(line)
	}

	checked = cac.containsAll(line, cac.items...)
	if !checked && cac.options.cache {
		checked = cac.containsAll(cac.all.String(), cac.items...)
	}

	return
}

func (cac *containsAllChecker) containsAll(content string, items ...string) (contains bool) {
	contains = true
	for _, item := range items {
		contains = contains && strings.Contains(content, item)
	}

	return
}
