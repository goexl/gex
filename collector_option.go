package gex

import (
	"strings"
)

type (
	collectorOption interface {
		applyCollector(options *collectorOptions)
	}

	collectorOptions struct {
		typ   OutputType
		mode  CollectorMode
		max   int
		lines []string
	}
)

func defaultCollectorOptions() *collectorOptions {
	return &collectorOptions{
		typ:   OutputTypeAny,
		mode:  CollectorModeCache,
		max:   1000,
		lines: make([]string, 0),
	}
}

func newCollectorOptions(opts ...collectorOption) (_options *collectorOptions) {
	_options = defaultCollectorOptions()
	for _, opt := range opts {
		opt.applyCollector(_options)
	}

	return
}

func (c *collectorOptions) write(line string) {
	c.lines = append(c.lines, line)
	current := len(c.lines)
	if current > c.max {
		c.lines = c.lines[current-c.max:]
	}
}

func (c *collectorOptions) string() string {
	return strings.Join(c.lines, ``)
}
