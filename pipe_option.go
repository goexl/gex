package gex

type (
	pipeOption interface {
		applyPipe(options *pipeOptions)
	}

	pipeOptions struct {
		args       []string
		dir        string
		envs       []*env
		systemEnvs bool
	}
)
