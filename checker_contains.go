package gex

import (
	`strings`
)

type containsChecker struct {
	contains string
}

func (c *containsChecker) check(all string, line string) (checked bool, err error) {
	checked = strings.Contains(line, c.contains)
	if !checked && `` != all {
		checked = strings.Contains(all, c.contains)
	}

	return
}
