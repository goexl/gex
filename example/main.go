package main

import (
	`github.com/storezhang/gex`
)

func main() {
	_, _ = gex.Exec(`ping`, gex.Args(`www.163.com`))
}
