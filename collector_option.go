package gex

import (
	`math`
	`strings`
)

type (
	collectorOption interface {
		applyCollector(options *collectorOptions)
	}

	collectorOptions struct {
		mode  CollectorMode
		max   int
		lines []string
	}
)

func defaultCollectorOptions() *collectorOptions {
	return &collectorOptions{
		mode:  CollectorModeAny,
		max:   math.MaxInt32,
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
	var sb strings.Builder
	for _, line := range c.lines {
		sb.WriteString(line)
	}

	return sb.String()
}
