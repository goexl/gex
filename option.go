package gex

import (
	`io`
	`os`
)

var _ = NewOptions

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

		async     bool
		wait      bool
		checker   checker
		collector collector
		charset   string
	}
)

// NewOptions 快捷方法，因为 option 接口不对外开放
func NewOptions(opts ...option) []option {
	return opts
}

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
