package gex

type counterParams struct {
	stream string
}

func newCounterParams() *counterParams {
	return new(counterParams)
}
