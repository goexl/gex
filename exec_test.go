package exec_test

import (
	`io/ioutil`
	`os`
	`testing`

	`github.com/golangex/exec`
)

func TestStartWithCheckerSuccess(t *testing.T) {
	t.Parallel()
	_, err := exec.Start(
		`ping`, exec.Args(`www.163.com`),
		exec.ContainsChecker(`Ping statistics for`), exec.Async(), exec.Quiet(),
	)
	if nil != err {
		t.FailNow()
	}
}

func TestStartWithCheckerFailed(t *testing.T) {
	t.Parallel()
	_, err := exec.Start(
		`ping`, exec.Args(`www.163.com`),
		exec.ContainsChecker(`xxx`), exec.Async(), exec.Quiet(),
	)
	if nil != err {
		t.FailNow()
	}
}

func TestStartWithSync(t *testing.T) {
	t.Parallel()
	_, err := exec.Start(`ping`, exec.Args(`www.163.com`), exec.Sync())
	if nil != err {
		t.FailNow()
	}
}

func TestStartWithStringCollector(t *testing.T) {
	t.Parallel()
	output := ``
	_, err := exec.Start(
		`ping`, exec.Args(`www.163.com`),
		exec.Sync(), exec.Quiet(), exec.StringCollector(&output),
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
	_, _ = exec.Start(`ping`, exec.Args(`abc.c`), exec.Quiet())

	_ = writer.Close()
	err, _ := ioutil.ReadAll(reader)
	os.Stderr = stderr

	if `` == string(err) {
		t.FailNow()
	}
}
