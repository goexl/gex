package gex

import (
	`io`
	`os`
)

var (
	_        = Stdout
	_        = FileStdout
	_        = FilenameStdout
	_ option = (*optionStdout)(nil)
)

type optionStdout struct {
	stdout io.Writer
	err    error
}

// Stdout 配置标准输出流
func Stdout(stdout io.Writer) *optionStdout {
	return &optionStdout{
		stdout: stdout,
	}
}

// FileStdout 文件输出
func FileStdout(file *os.File) *optionStdout {
	return &optionStdout{
		stdout: file,
	}
}

// FilenameStdout 文件名输出
func FilenameStdout(filename string, opts ...fileOption) (stdout *optionStdout) {
	stdout = new(optionStdout)
	if file, err := parseFile(filename, opts...); nil != err {
		stdout.err = err
	} else {
		stdout.stdout = file
	}

	return
}

func (s *optionStdout) apply(options *options) (err error) {
	if err = s.err; nil != err {
		return
	}
	options.collectors[keyStdout] = &writerCollector{
		writer: s.stdout,
		options: &collectorOptions{
			typ: OutputTypeStdout,
		},
	}

	return
}
