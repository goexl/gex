package gex

var _ collector = (*stringsCollector)(nil)

type stringsCollector struct {
	strings *[]string
	params  *collectorParams
}

func newStringsCollector(strings *[]string, params *collectorParams) *stringsCollector {
	return &stringsCollector{
		strings: strings,
		params:  params,
	}
}

func (sc *stringsCollector) Collect(line string, stream string) (err error) {
	if "" == sc.params.stream || stream == sc.params.stream {
		line = sc.params.process(line)
		*sc.strings = append(*sc.strings, line)
	}

	return
}

func (sc *stringsCollector) Notify(_ int, _ error) {}
