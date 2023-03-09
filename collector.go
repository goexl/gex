package gex

type collector interface {
	Collect(line string, stream string) (err error)
}
