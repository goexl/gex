package guide

import (
	"github.com/goexl/gex"
	"github.com/goexl/gox/args"
)

var _ = cli

func cli() (int, error) {
	return gex.New("ping").
		Args(args.New().Build().Add("www.163.com").Flag("c").Add(10).Build()).
		Build().
		Exec()
}
