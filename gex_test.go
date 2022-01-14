package gex_test

import (
	`testing`

	`github.com/storezhang/gex`
)

func TestRun(t *testing.T) {
	_, err := gex.Run(
		`D:\Downloads\fastgithub_win-x64\fastgithub.exe`,
		gex.ContainsChecker(`127.0.0.1:7890`), gex.Async(), gex.Quiet(),
	)
	if nil != err {
		t.FailNow()
	}
}
