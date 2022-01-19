package gex

var _ = NewOptions

type (
	option interface {
		apply(options *options)
	}

	options struct {
		args       []string
		dir        string
		envs       []*env
		systemEnvs bool

		async      bool
		wait       bool
		checker    checker
		collectors map[string]collector
		charset    string

		notifiers []notifier
	}
)

// NewOptions 快捷方法，因为接口不对外开放
func NewOptions(opts ...option) []option {
	return opts
}

func defaultOptions() *options {
	return &options{
		collectors: map[string]collector{
			keyTerminal: newTerminalCollector(),
		},

		systemEnvs: true,
		async:      false,
	}
}
