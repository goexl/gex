package guide

import (
	"github.com/goexl/gex"
)

var _ = dir

func dir() (int, error) {
	return gex.Exec(`redis`, gex.Dir(`/opt/`))
}
