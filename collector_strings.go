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
		*sc.strings = append(*sc.strings, line)
	}

	return
}

func (sc *stringsCollector) Name() string {
	return "strings"
}

func (sc *stringsCollector) Notify(_ int, _ error) {}
