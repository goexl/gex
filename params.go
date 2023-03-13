package gex

import (
	"context"
	"io"

	"github.com/goexl/gox/args"
)

type params struct {
	context     context.Context
	name        string
	args        *args.Args
	dir         string
	environment *environmentParams

	async   bool
	pwe     bool
	echo    bool
	stdin   io.Reader
	charset string

	checks     []*check
	collectors map[string]collector
	notifiers  []notifier
	counters   []counter
}

func newParams(name string) *params {
	return &params{
		name: name,

		async:      false,
		checks:     make([]*check, 0),
		collectors: make(map[string]collector),
		notifiers:  make([]notifier, 0),
		counters:   make([]counter, 0),
	}
}

func (p *params) check(line string) (checked bool, err error) {
	checked = true
	if 0 == len(p.checks) {
		return
	}

	size := len(p.checks)
	for index, _check := range p.checks {
		if result, ce := _check.checker.Check(line); nil != ce {
			err = ce
		} else if operatorAnd == _check.operator {
			checked = checked && result
		} else if operatorOr == _check.operator {
			checked = checked || result
		}
		if nil != err || index == size-1 {
			break
		}

		// 实现逻辑断路器
		next := p.checks[index+1]
		if operatorAnd == next.operator && !checked || operatorOr == next.operator && checked {
			break
		}
	}

	return
}
