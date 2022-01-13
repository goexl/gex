package gex

import (
	`io`
	`os`
)

type (
	option interface {
		apply(options *options)
	}

	options struct {
		args       []string
		outs       []io.Writer
		errs       []io.Writer
		dir        string
		envs       []*env
		systemEnvs bool

		async   bool
		checker checker
	}
)

func defaultOptions() *options {
	return &options{
		outs: []io.Writer{
			os.Stdout,
		},
		errs: []io.Writer{
			os.Stderr,
		},

		systemEnvs: true,
		async:      false,
	}
}
