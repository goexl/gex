package main

import (
	`github.com/goexl/gex`
)

func main() {
	_, _ = gex.Exec(`ping`, gex.Args(`www.163.com`, `-c`, 10))
}
