package gex

var _ = New

// New 合建命令
func New(command string) *Builder {
	return newBuilder(command)
}
