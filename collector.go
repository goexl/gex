package gex

type collector interface {
	key() string
	collect(line string, mode CollectorMode) (err error)
}
