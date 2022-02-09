package gex

type collector interface {
	Collect(line string, ot OutputType) (err error)
}
