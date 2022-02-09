package gex

var (
	_        = Notifier
	_        = FuncNotifier
	_ option = (*optionNotifier)(nil)
)

type optionNotifier struct {
	notifier notifier
}

// Notifier 配置检查器
func Notifier(notifier notifier) *optionNotifier {
	return &optionNotifier{
		notifier: notifier,
	}
}

// FuncNotifier 应用内方法通知
func FuncNotifier(fun notifyFunc) *optionNotifier {
	return &optionNotifier{
		notifier: &funcNotifier{
			fun: fun,
		},
	}
}

func (n *optionNotifier) apply(options *options) {
	options.notifiers = append(options.notifiers, n.notifier)
}
