package gex

import (
	`math`
)

type (
	collectorOption interface {
		applyCollector(options *collectorOptions)
	}

	collectorOptions struct {
		mode    CollectorMode
		max     int64
		current int64
	}
)

func defaultCollectorOptions() *collectorOptions {
	return &collectorOptions{
		mode: CollectorModeAny,
		max:  math.MaxInt64,
	}
}

func (c *collectorOptions) canCollect() bool {
	return c.current < c.max
}
