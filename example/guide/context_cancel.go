package guide

import (
	"context"
	"time"

	"github.com/goexl/gex"
)

var _ = cancel

func cancel() (code int, err error) {
	ctx, _cancel := context.WithCancel(context.Background())
	code, err = gex.New("redis").Context(ctx).Build().Exec()
	time.Sleep(time.Minute)
	_cancel()

	return
}
