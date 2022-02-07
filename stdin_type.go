package gex

const (
	stdinTypeUnknown stdinType = 0
	stdinTypePipe    stdinType = 1
	stdinTypeFile    stdinType = 2
)

type stdinType uint8
