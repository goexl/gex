package gex_test

import (
	`io/ioutil`
	`os`
	`testing`

	`github.com/goexl/gex`
)

func TestStartWithCheckerSuccess(t *testing.T) {
	t.Parallel()
	_, err := gex.Exec(
		`ping`, gex.Args(`www.163.com`),
		gex.ContainsChecker(`Ping statistics for`), gex.Async(), gex.Quiet(),
	)
	if nil != err {
		t.FailNow()
	}
}

func TestStartWithCheckerFailed(t *testing.T) {
	t.Parallel()
	_, err := gex.Exec(
		`ping`, gex.Args(`www.163.com`),
		gex.ContainsChecker(`xxx`), gex.Async(), gex.Quiet(),
	)
	if nil != err {
		t.FailNow()
	}
}

func TestStartWithSync(t *testing.T) {
	t.Parallel()
	_, err := gex.Exec(`ping`, gex.Args(`www.163.com`), gex.Sync())
	if nil != err {
		t.FailNow()
	}
}

func TestStartWithStringCollector(t *testing.T) {
	t.Parallel()
	output := ``
	_, err := gex.Exec(
		`ping`, gex.Args(`www.163.com`),
		gex.Sync(), gex.Quiet(), gex.StringCollector(&output),
	)
	if nil != err || `` == output {
		t.FailNow()
	}
}

func TestStartPwe(t *testing.T) {
	t.Parallel()
	stderr := os.Stderr
	reader, writer, _ := os.Pipe()
	os.Stderr = writer

	// 执行命令，该命令一定是错误的
	_, _ = gex.Exec(`ping`, gex.Args(`abc.c`), gex.Quiet())

	_ = writer.Close()
	err, _ := ioutil.ReadAll(reader)
	os.Stderr = stderr

	if `` == string(err) {
		t.FailNow()
	}
}
