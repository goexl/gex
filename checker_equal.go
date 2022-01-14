package gex

import (
	`strings`
)

type equalChecker struct {
	equal string

	cache bool
	all   strings.Builder
}

func (ec *equalChecker) check(line string) (checked bool, err error) {
	if ec.cache {
		ec.all.WriteString(line)
	}

	checked = ec.equal == line
	if !checked && ec.cache {
		checked = ec.all.String() == ec.equal
	}

	return
}
