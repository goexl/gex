package gex

import (
	`io`
)

type (
	option interface {
		apply(options *options)
	}

	options struct {
		args       []string
		stdout     io.Writer
		stderr     io.Writer
		dir        string
		envs       []*env
		systemEnvs bool

		async bool
	}
)

func defaultOptions() *options {
	return &options{
		systemEnvs: true,
		async:      false,
	}
}
