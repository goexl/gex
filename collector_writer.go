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
			mode: CollectorModeAny,
		},
	}
}

func (w *writerCollector) collect(line string, mode CollectorMode) (err error) {
	if CollectorModeAny != w.options.mode && mode != w.options.mode {
		return
	}
	w.options.write(line)

	return
}

func (w *writerCollector) notify(_ int, _ error) {
	_, _ = w.writer.Write([]byte(w.options.string()))
}
