package param

import (
	"context"
	"io"

	"github.com/goexl/args"
	"github.com/goexl/gex/internal/internal/core"
)

type Command struct {
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

	Logics     []*core.Logic
	Collectors map[string]core.Collector
	Notifiers  []core.Notifier
	Counters   []core.Counter
}

func NewCommand(name string) *Command {
	return &Command{
		Name: name,

		Async:      false,
		Logics:     make([]*core.Logic, 0),
		Collectors: make(map[string]core.Collector),
		Notifiers:  make([]core.Notifier, 0),
		Counters:   make([]core.Counter, 0),
	}
}

func (c *Command) Check(line string) (checked bool, err error) {
	checked = true
	if 0 == len(c.Logics) {
		return
	}

	size := len(c.Logics)
	for index, logic := range c.Logics {
		checked, err = logic.Check(checked, line)
		if index == size-1 {
			break
		} else if !c.Logics[index+1].Next(checked) {
			break
		}
	}

	return
}
