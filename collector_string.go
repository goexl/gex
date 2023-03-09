package gex

import (
	"strings"
)

var _ collector = (*stringCollector)(nil)

type stringCollector struct {
	string *string
	params *collectorParams

	lines []string
}

func newTerminalCollector(output *string) *stringCollector {
	return newStringCollector(output, &collectorParams{
		max: 1024,
	})
}

func newStringCollector(output *string, params *collectorParams) *stringCollector {
	return &stringCollector{
		string: output,
		params: params,

		lines: make([]string, 0, params.max),
	}
}

func (sc *stringCollector) Collect(line string, stream string) (err error) {
	if "" != sc.params.stream && stream != sc.params.stream {
		return
	}

	line = sc.params.process(line)
	sc.lines = append(sc.lines, line)
	current := len(sc.lines)
	if current > sc.params.max {
		sc.lines = sc.lines[current-sc.params.max:]
	}

	return
}

func (sc *stringCollector) Notify(_ int, _ error) {
	*sc.string = strings.Join(sc.lines, empty)
}
