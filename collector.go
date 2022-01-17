package gex

type collector interface {
	collect(line string)
}
