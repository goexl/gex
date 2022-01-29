package gex

type collector interface {
	notifier

	collect(line string, typ CollectorType) (err error)
}
