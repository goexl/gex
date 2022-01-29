package gex

import (
	`strings`
)

type (
	collectorOption interface {
		applyCollector(options *collectorOptions)
	}

	collectorOptions struct {
		typ   CollectorType
		mode  CollectorMode
		max   int
		lines []string
	}
)

func defaultCollectorOptions() *collectorOptions {
	return &collectorOptions{
		typ:   CollectorTypeAny,
		mode:  CollectorModeCache,
		max:   1000,
		lines: make([]string, 0),
	}
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
