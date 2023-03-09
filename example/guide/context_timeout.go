package guide

import (
	"context"
	"time"

	"github.com/goexl/gex"
)

var _ = timeout

func timeout() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	return gex.New("redis").Context(ctx).Build().Exec()
}
