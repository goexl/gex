package gex

import (
	`io`
	`os`
)

var (
	_        = Stderr
	_        = FileStderr
	_        = FilenameStderr
	_ option = (*optionStderr)(nil)
)

type optionStderr struct {
	stderr io.Writer
	err    error
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

// FilenameStderr 文件名输出
func FilenameStderr(filename string, opts ...fileOption) (stderr *optionStderr) {
	stderr = new(optionStderr)
	if file, err := parseFile(filename, opts...); nil != err {
		stderr.err = err
	} else {
		stderr.stderr = file
	}

	return
}

func (s *optionStderr) apply(options *options) (err error) {
	if err = s.err; nil != err {
		return
	}
	options.collectors[keyStderr] = &writerCollector{
		writer: s.stderr,
		options: &collectorOptions{
			typ: OutputTypeStderr,
		},
	}

	return
}
