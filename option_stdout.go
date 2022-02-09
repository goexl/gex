package gex

import (
	`io`
	`os`
)

var (
	_        = Stdout
	_        = FileStdout
	_ option = (*optionStdout)(nil)
)

type optionStdout struct {
	stdout io.Writer
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

func (s *optionStdout) apply(options *options) {
	options.collectors[keyStdout] = &writerCollector{
		writer: s.stdout,
		options: &collectorOptions{
			typ: OutputTypeStdout,
		},
	}
}
