package pipe

import (
	`context`
	`time`

	`github.com/storezhang/gex`
)

var _ = cancel

func cancel() (code int, err error) {
	ctx, _cancel := context.WithCancel(context.Background())
	code, err = gex.Exec(`ping`, gex.Pipe(`echo`, gex.Args(`www.163.con`), gex.Context(ctx)))
	time.Sleep(time.Minute)
	_cancel()

	return
}
