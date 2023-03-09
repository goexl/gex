package gex

type collectorParams struct {
	stream string
	max    int
	trim   *trim
}

func newCollectorParams() *collectorParams {
	return &collectorParams{
		max:  1000,
		trim: newTrim(),
	}
}
