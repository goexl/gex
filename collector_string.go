package exec

var _ collector = (*stringCollector)(nil)

type stringCollector struct {
	string  *string
	options *collectorOptions
}

func newOutputStringCollector(output *string, max int) *stringCollector {
	return &stringCollector{
		string: output,
		options: &collectorOptions{
			typ:  OutputTypeAny,
			mode: CollectorModeCache,
			max:  max,
		},
	}
}

func (s *stringCollector) Collect(line string, ot OutputType) (err error) {
	if OutputTypeAny != s.options.typ && ot != s.options.typ {
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

func (s *stringCollector) Notify(_ int, _ error) {
	if CollectorModeCache == s.options.mode {
		*s.string = s.options.string()
	}
}
