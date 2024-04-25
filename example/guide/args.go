package guide

import (
	"github.com/goexl/args"
	"github.com/goexl/gex"
)

var _ = cli

func cli() (int, error) {
	return gex.New("ping").
		Args(args.New().Build().Add("www.163.com").Flag("c").Add(10).Build()).
		Build().
		Exec()
}
