package gex

type notifier interface {
	notify(code int, err error)
}
