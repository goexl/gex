package pipe

import (
	`context`
	`time`

	`github.com/goexl/gex`
)

var _ = timeout

func timeout() (int, error) {
	ctx, _cancel := context.WithTimeout(context.Background(), time.Minute)
	defer _cancel()

	return gex.Exec(`redis`, gex.Pipe(`echo`, gex.Args(`www.163.con`), gex.Context(ctx)))
}
