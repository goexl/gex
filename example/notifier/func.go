package notifier

import (
	`fmt`

	`github.com/goexl/gex`
)

var _ = fun

func fun() (int, error) {
	return gex.Exec(`redis`, gex.FuncNotifier(notify))
}

func notify(code int, err error) {
	fmt.Println(code, err)
}
