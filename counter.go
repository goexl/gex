package gex

type counter interface {
	// Count 计数
	Count(line string, ot OutputType) error
}
