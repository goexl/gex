package gex

type collector interface {
	notifier

	key() string
	collect(line string, mode CollectorMode) (err error)
}
