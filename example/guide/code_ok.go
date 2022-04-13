package guide

import (
	"github.com/goexl/gex"
)

var _ = codeOk

func codeOk() (int, error) {
	return gex.Exec(`redis`, gex.OkCode(1217))
}
