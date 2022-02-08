package gex

var _ collector = (*stringCollector)(nil)

type stringCollector struct {
	string  *string
	options *collectorOptions
}

func newOutputStringCollector(output *string, max int) *stringCollector {
	return &stringCollector{
		string: output,
		options: &collectorOptions{
			typ: CollectorTypeAny,
			max: max,
		},
	}
}

func (s *stringCollector) Collect(line string, typ CollectorType) (err error) {
	if CollectorTypeAny != s.options.typ && typ != s.options.typ {
		return
	}

	switch s.options.mode {
	case CollectorModeDirect:
		*s.string = line
	default:
		s.options.write(line)
	}

	return
}

func (s *stringCollector) notify(_ int, _ error) {
	if CollectorModeCache == s.options.mode {
		*s.string = s.options.string()
	}
}
