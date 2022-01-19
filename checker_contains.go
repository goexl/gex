package gex

import (
	`strings`
)

type containsChecker struct {
	contains string

	options *checkerOptions
	all     strings.Builder
}

func (cc *containsChecker) check(line string) (checked bool, err error) {
	if cc.options.cache {
		cc.all.WriteString(line)
	}

	checked = strings.Contains(line, cc.contains)
	if !checked && cc.options.cache {
		checked = strings.Contains(cc.all.String(), cc.contains)
	}

	return
}
