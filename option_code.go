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
func Code(ok int) *optionCode {
	return &optionCode{
		code: &code{
			ok: ok,
		},
	}
}

// OkCode 正常退出状态码
func OkCode(ok int) *optionCode {
	return &optionCode{
		code: &code{
			ok: ok,
		},
	}
}

func (c *optionCode) apply(options *options) (err error) {
	options.code = c.code

	return
}
