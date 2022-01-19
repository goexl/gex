package gex

import (
	`fmt`
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

func (w *writerCollector) key() string {
	return fmt.Sprintf(`%p`, w.writer)
}

func (w *writerCollector) collect(line string, mode CollectorMode) (err error) {
	if CollectorModeAny != w.options.mode && mode != w.options.mode {
		return
	}
	_, err = w.writer.Write([]byte(line))

	return
}
