package gex

import (
	`fmt`
)

var _ collector = (*stringCollector)(nil)

type stringCollector struct {
	string  *string
	options *collectorOptions
}

func newOutputStringCollector(output *string, max int) *stringCollector {
	return &stringCollector{
		string: output,
		options: &collectorOptions{
			mode: CollectorModeAny,
			max:  max,
		},
	}
}

func (s *stringCollector) key() string {
	return fmt.Sprintf(`%p`, s.string)
}

func (s *stringCollector) collect(line string, mode CollectorMode) (err error) {
	if CollectorModeAny != s.options.mode && mode != s.options.mode {
		return
	}
	s.options.write(line)

	return
}

func (s *stringCollector) notify(_ int, _ error) {
	*s.string = s.options.string()
}
