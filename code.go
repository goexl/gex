package gex

type code struct {
	ok int
}

// NewCode 创建代码
func NewCode(ok int) *code {
	return &code{
		ok: ok,
	}
}
