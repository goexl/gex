package gex

type (
	counterOption interface {
		applyCounter(options *counterOptions)
	}

	counterOptions struct {
		typ OutputType
	}
)

func defaultCounterOptions() *counterOptions {
	return &counterOptions{
		typ: OutputTypeAny,
	}
}

func newCounterOptions(opts ...counterOption) (_options *counterOptions) {
	_options = defaultCounterOptions()
	for _, opt := range opts {
		opt.applyCounter(_options)
	}

	return
}
