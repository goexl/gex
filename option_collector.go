package gex

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var (
	_        = Collector
	_        = StringCollector
	_        = FileCollector
	_        = FilenameCollector
	_        = WriterCollector
	_ option = (*optionCollector)(nil)
)

type optionCollector struct {
	collector collector
	err       error
}

// Collector 配置输出
func Collector(collector collector) *optionCollector {
	return &optionCollector{
		collector: collector,
	}
}

// StringCollector 配置输出到字符串
func StringCollector(string *string, opts ...collectorOption) *optionCollector {
	return &optionCollector{
		collector: &stringCollector{
			string:  string,
			options: newCollectorOptions(opts...),
		},
	}
}

// FileCollector 配置输出到文件
func FileCollector(file *os.File, opts ...collectorOption) *optionCollector {
	return &optionCollector{
		collector: &writerCollector{
			writer:  file,
			options: newCollectorOptions(opts...),
		},
	}
}

// WriterCollector 配置输出到文件
func WriterCollector(writer io.Writer, opts ...collectorOption) *optionCollector {
	return &optionCollector{
		collector: &writerCollector{
			writer:  writer,
			options: newCollectorOptions(opts...),
		},
	}
}

// FilenameCollector 配置输出到文件
func FilenameCollector(filename string, opts ...fileOption) (collector *optionCollector) {
	collector = new(optionCollector)
	if file, err := parseFile(filename, opts...); nil != err {
		collector.err = err
	} else {
		collector.collector = &writerCollector{
			writer: bufio.NewWriter(file),
		}
	}

	return
}

func (c *optionCollector) apply(options *options) (err error) {
	if err = c.err; nil != err {
		return
	}
	options.collectors[fmt.Sprintf(`%p`, c.collector)] = c.collector

	return
}
