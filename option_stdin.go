package gex

import (
	`io`
	`os`
)

var (
	_        = ReaderStdin
	_        = FileStdin
	_ option = (*optionStdin)(nil)
)

type optionStdin struct {
	reader io.Reader
}

// ReaderStdin 读取者输入
func ReaderStdin(reader io.Reader) *optionStdin {
	return &optionStdin{
		reader: reader,
	}
}

// FileStdin 文件输入
func FileStdin(file *os.File) *optionStdin {
	return &optionStdin{
		reader: file,
	}
}

func (s *optionStdin) apply(options *options) {
	options.stdin = s.reader
}
