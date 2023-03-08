package gex

var _ collector = (*stringsCollector)(nil)

type stringsCollector struct {
	strings *[]string
	options *collectorOptions
}

func (s *stringsCollector) Collect(line string, ot OutputType) (err error) {
	if OutputTypeAny != s.options.typ && ot != s.options.typ {
		return
	}
	*s.strings = append(*s.strings, line)

	return
}

func (s *stringsCollector) Notify(_ int, _ error) {}
