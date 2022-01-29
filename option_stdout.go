package gex

import (
	`io`
)

var (
	_        = Stdout
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

func (s *optionStdout) apply(options *options) {
	options.collectors[keyStdout] = &writerCollector{
		writer: s.stdout,
		options: &collectorOptions{
			typ: CollectorTypeStdout,
		},
	}
}
