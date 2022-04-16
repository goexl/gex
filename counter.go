package gex

type counter interface {
	// Count 计数
	Count(line string, typ OutputType) error
}
