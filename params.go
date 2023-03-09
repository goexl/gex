package gex

import (
	"context"
	"io"

	"github.com/goexl/gox/args"
)

type params struct {
	context context.Context
	name    string
	args    *args.Args
	dir     string
	envs    []string
	system  bool

	async bool
	stdin io.Reader

	checker    checker
	collectors map[string]collector
	charset    string

	notifiers []notifier
	counters  []counter

	max int
}

func newParams(name string) *params {
	return &params{
		name: name,
		collectors: map[string]collector{
			keyTerminal: newTerminalCollector(),
		},
		async: false,
		max:   100,
	}
}
