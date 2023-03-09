package guide

import (
	"github.com/goexl/gex"
)

var _ = envAdd

func envAdd() (int, error) {
	return gex.Exec(`redis`, gex.Env(`USERNAME`, `storezhang`))
}
