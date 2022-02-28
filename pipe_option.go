package exec

import (
	`context`
)

type (
	pipeOption interface {
		applyPipe(options *pipeOptions)
	}

	pipeOptions struct {
		context    context.Context
		args       []string
		dir        string
		envs       []*env
		systemEnvs bool
	}
)

func defaultPipeOptions() *pipeOptions {
	return &pipeOptions{
		systemEnvs: true,
	}
}
