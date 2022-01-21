package gex

type collector interface {
	notifier

	collect(line string, mode CollectorMode) (err error)
}
