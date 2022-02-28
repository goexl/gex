package exec

import (
	`strings`
)

type equalChecker struct {
	equal string

	options *checkerOptions
	all     strings.Builder
}

func (ec *equalChecker) Check(line string) (checked bool, err error) {
	if ec.options.cache {
		ec.all.WriteString(line)
	}

	checked = ec.equal == line
	if !checked && ec.options.cache {
		checked = ec.all.String() == ec.equal
	}

	return
}
