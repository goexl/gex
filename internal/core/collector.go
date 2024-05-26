package core

type Collector interface {
	Collect(line string, stream string) (err error)
}
