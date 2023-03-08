package gex

import (
	"context"
	"io"

	"github.com/goexl/gox"
)

type params struct {
	context context.Context
	args    *gox.Args
	dir     string
	envs    []string
	system  *system
	code    *code

	async bool

	stdin io.Reader
	pipes []*pipe

	checker    checker
	collectors map[string]collector
	charset    string

	notifiers []notifier
	counters  []counter

	pwe bool
	max int
}

func newParams() *params {
	return &params{
		collectors: map[string]collector{
			keyTerminal: newTerminalCollector(),
		},

		system: &system{
			envs: true,
		},
		code: &code{
			ok: defaultCodeOk,
		},
		async: false,

		pwe: true,
		max: 100,
	}
}
