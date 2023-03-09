package gex

type collectorParams struct {
	stream string
	max    int
	trim   *trim
}

func newCollectorParams() *collectorParams {
	return &collectorParams{
		max:  1024,
		trim: newTrim(),
	}
}

func (cp *collectorParams) process(line string) (final string) {
	final = line
	if nil != cp.trim {
		final = cp.trim.process(line)
	}

	return
}
