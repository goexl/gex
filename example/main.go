package main

import (
	"github.com/goexl/args"
	"github.com/goexl/gex"
)

func main() {
	_, _ = gex.New("ping").Args(args.New().Build().Subcommand("www.163.com").Flag("-c").Add(10).Build()).Build().Exec()
}
