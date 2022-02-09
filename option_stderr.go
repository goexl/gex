package gex

import (
	`io`
	`os`
)

var (
	_        = Stderr
	_        = FileStderr
	_ option = (*optionStderr)(nil)
)

type optionStderr struct {
	stderr io.Writer
}

// Stderr 配置标准错误流
func Stderr(stderr io.Writer) *optionStderr {
	return &optionStderr{
		stderr: stderr,
	}
}

// FileStderr 文件输出
func FileStderr(file *os.File) *optionStderr {
	return &optionStderr{
		stderr: file,
	}
}

func (s *optionStderr) apply(options *options) {
	options.collectors[keyStderr] = &writerCollector{
		writer: s.stderr,
		options: &collectorOptions{
			typ: OutputTypeStderr,
		},
	}
}
