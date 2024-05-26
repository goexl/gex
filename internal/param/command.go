package param

import (
	"context"
	"io"

	"github.com/goexl/args"
	"github.com/goexl/gex/internal/core"
	core2 "github.com/goexl/gex/internal/internal/core"
	"github.com/goexl/gox"
)

type Command struct {
	gox.CannotCopy

	Context     context.Context
	Name        string
	Arguments   *args.Arguments
	Directory   string
	Environment *Environment

	Async   bool
	Pwe     bool
	Echo    bool
	Stdin   io.Reader
	Charset string

	Checks     []*core2.Check
	Collectors map[string]core.Collector
	Notifiers  []core.Notifier
	Counters   []core.Counter
}

func NewCommand(name string) *Command {
	return &Command{
		Name: name,

		Async:      false,
		Checks:     make([]*core2.Check, 0),
		Collectors: make(map[string]core.Collector),
		Notifiers:  make([]core.Notifier, 0),
		Counters:   make([]core.Counter, 0),
	}
}

func (c *Command) Check(line string) (checked bool, err error) {
	checked = true
	if 0 == len(c.Checks) {
		return
	}

	size := len(c.Checks)
	for index, check := range c.Checks {
		if result, ce := check.Checker.Check(line); nil != ce {
			err = ce
		} else if core.OperatorAnd == check.Operator {
			checked = checked && result
		} else if core.OperatorOr == check.Operator {
			checked = checked || result
		}
		if nil != err || index == size-1 {
			break
		}

		// 实现逻辑断路器
		next := c.Checks[index+1]
		if core.OperatorAnd == next.Operator && !checked || core.OperatorOr == next.Operator && checked {
			break
		}
	}

	return
}
