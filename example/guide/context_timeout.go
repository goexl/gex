package guide

import (
	`context`
	`time`

	`github.com/storezhang/gex`
)

var _ = timeout

func timeout() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	return gex.Exec(`redis`, gex.Context(ctx))
}
