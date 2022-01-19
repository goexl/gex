package gex

import (
	`fmt`
	`strings`
)

var _ collector = (*stringCollector)(nil)

type stringCollector struct {
	string  *string
	options *collectorOptions
}

func (s *stringCollector) key() string {
	return fmt.Sprintf(`%p`, s.string)
}

func (s *stringCollector) collect(line string, mode CollectorMode) (err error) {
	if CollectorModeAny != s.options.mode && mode != s.options.mode {
		return
	}

	var sb strings.Builder
	sb.WriteString(*s.string)
	sb.WriteString(line)
	*s.string = sb.String()

	return
}
