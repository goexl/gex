package notifier

import (
	`fmt`

	`github.com/goexl/gex`
)

var _ = usage

type notifier struct{}

func newNotifier() *notifier {
	return new(notifier)
}

func (c *notifier) Notify(code int, err error) {
	fmt.Println(code, err)
}

func usage() (int, error) {
	return gex.Exec(`redis`, gex.Notifier(newNotifier()))
}
