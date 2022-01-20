package gex

import (
	`bufio`
	`os`
)

var (
	_        = Collector
	_        = StringCollector
	_        = FileCollector
	_        = FilenameCollector
	_ option = (*optionCollector)(nil)
)

type optionCollector struct {
	collector collector
}

// Collector 配置输出
func Collector(collector collector) *optionCollector {
	return &optionCollector{
		collector: collector,
	}
}

// StringCollector 配置输出到字符串
func StringCollector(string *string, opts ...collectorOption) *optionCollector {
	_options := defaultCollectorOptions()
	for _, opt := range opts {
		opt.applyCollector(_options)
	}

	return &optionCollector{
		collector: &stringCollector{
			string:  string,
			options: _options,
		},
	}
}

// FileCollector 配置输出到文件
func FileCollector(file *os.File, opts ...collectorOption) *optionCollector {
	_options := defaultCollectorOptions()
	for _, opt := range opts {
		opt.applyCollector(_options)
	}

	return &optionCollector{
		collector: &writerCollector{
			writer:  bufio.NewWriter(file),
			options: _options,
		},
	}
}

// FilenameCollector 配置输出到文件
func FilenameCollector(filename string, opts ...fileOption) (collector *optionCollector) {
	file, err := parseFile(filename, opts...)
	if nil == err {
		panic(err)
	} else if nil != file {
		collector = new(optionCollector)
		collector.collector = &writerCollector{
			writer: bufio.NewWriter(file),
		}
	}

	return
}

func (c *optionCollector) apply(options *options) {
	options.collectors[c.collector.key()] = c.collector
}
