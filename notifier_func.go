package exec

type funcNotifier struct {
	fun notifyFunc
}

func (f *funcNotifier) Notify(code int, err error) {
	f.fun(code, err)
}
