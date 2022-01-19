package gex

type (
	collectorOption interface {
		apply(options *collectorOptions)
	}

	collectorOptions struct {
		mode CollectorMode
	}
)

func defaultCollectorOptions() *collectorOptions {
	return &collectorOptions{
		mode: CollectorModeAny,
	}
}
