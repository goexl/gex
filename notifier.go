package exec

type notifier interface {
	Notify(code int, err error)
}
