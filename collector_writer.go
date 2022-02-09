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
			typ:  OutputTypeAny,
		},
	}
}

func (w *writerCollector) Collect(line string, ot OutputType) (err error) {
	if OutputTypeAny != w.options.typ && ot != w.options.typ {
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

func (w *writerCollector) Notify(_ int, _ error) {
	if CollectorModeCache == w.options.mode {
		_, _ = w.writer.Write([]byte(w.options.string()))
	}
}
