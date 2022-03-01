package gex

type notifier interface {
	Notify(code int, err error)
}
