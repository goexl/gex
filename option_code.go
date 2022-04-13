package gex

var (
	_        = Code
	_        = OkCode
	_ option = (*optionCode)(nil)
)

type optionCode struct {
	code *code
}

// Code 配置状态码
func Code(code *code) *optionCode {
	return &optionCode{
		code: code,
	}
}

// OkCode 正常退出状态码
func OkCode(c int) *optionCode {
	return Code(&code{
		ok: c,
	})
}

func (c *optionCode) apply(options *options) (err error) {
	options.code = c.code

	return
}
