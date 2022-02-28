package exec

import (
	`strings`
)

type containsAnyChecker struct {
	items []string

	options *checkerOptions
	all     strings.Builder
}

func (cac *containsAnyChecker) Check(line string) (checked bool, err error) {
	if cac.options.cache {
		cac.all.WriteString(line)
	}

	checked = cac.containsAny(line, cac.items...)
	if !checked && cac.options.cache {
		checked = cac.containsAny(cac.all.String(), cac.items...)
	}

	return
}

func (cac *containsAnyChecker) containsAny(content string, items ...string) (contains bool) {
	for _, item := range items {
		contains = strings.Contains(content, item)
		if contains {
			break
		}
	}

	return
}
