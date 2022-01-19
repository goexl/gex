package gex_test

import (
	`testing`

	`github.com/storezhang/gex`
)

func TestRunWithChecker(t *testing.T) {
	_, err := gex.Run(
		`ping`, gex.Args(`www.163.com`, `-n`, `1`),
		gex.ContainsChecker(`Ping statistics for`), gex.Async(), gex.Quiet(),
	)
	if nil != err {
		t.FailNow()
	}
}

func TestRunWithSync(t *testing.T) {
	_, err := gex.Run(`ping`, gex.Args(`www.163.com`, `-n`, `1`), gex.Sync())
	if nil != err {
		t.FailNow()
	}
}

func TestRunWithStringCollector(t *testing.T) {
	output := ``
	_, err := gex.Run(
		`ping`, gex.Args(`www.163.com`, `-n`, `1`),
		gex.Sync(), gex.Quiet(), gex.StringCollector(&output),
	)
	if nil != err || `` == output {
		t.FailNow()
	}
}
