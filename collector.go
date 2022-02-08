package gex

type collector interface {
	notifier

	Collect(line string, typ CollectorType) (err error)
}
