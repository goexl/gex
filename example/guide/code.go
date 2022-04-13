package guide

import (
	"github.com/goexl/gex"
)

var _ = code

func code() (int, error) {
	return gex.Exec(`redis`, gex.Code(1217))
}
