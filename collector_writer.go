package gex

import (
	`io`
	`os`
)

var _ collector = (*writerCollector)(nil)

type writerCollector struct {
	writer  io.Writer
	options *collectorOptions
}

func newTerminalCollector() *writerCollector {
	return &writerCollector{
		writer: os.Stdout,
		options: &collectorOptions{
			mode: CollectorModeDirect,
			typ:  CollectorTypeAny,
		},
	}
}

func (w *writerCollector) Collect(line string, typ CollectorType) (err error) {
	if CollectorTypeAny != w.options.typ && typ != w.options.typ {
		return
	}

	switch w.options.mode {
	case CollectorModeDirect:
		_, _ = w.writer.Write([]byte(line))
	default:
		w.options.write(line)
	}

	return
}

func (w *writerCollector) notify(_ int, _ error) {
	if CollectorModeCache == w.options.mode {
		_, _ = w.writer.Write([]byte(w.options.string()))
	}
}
