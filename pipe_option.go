package gex

import (
	"context"
)

type (
	pipeOption interface {
		applyPipe(options *pipeOptions)
	}

	pipeOptions struct {
		context context.Context
		args    []string
		dir     string
		envs    []*env
		system  *system
	}
)

func defaultPipeOptions() *pipeOptions {
	return &pipeOptions{
		system: &system{
			envs: true,
		},
	}
}
