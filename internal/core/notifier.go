package core

type Notifier interface {
	Notify(code int, err error)
}
