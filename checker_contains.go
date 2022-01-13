package gex

import (
	`strings`
)

type containsChecker struct {
	contains string
}

func (cc *containsChecker) check(all string, line string) (checked bool, err error) {
	checked = strings.Contains(line, cc.contains)
	if !checked && `` != all {
		checked = strings.Contains(all, cc.contains)
	}

	return
}
