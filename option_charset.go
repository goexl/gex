package gex

var (
	_        = Charset
	_ option = (*optionCharset)(nil)
)

type optionCharset struct {
	charset string
}

// Charset 配置流字符集
func Charset(charset string) *optionCharset {
	return &optionCharset{
		charset: charset,
	}
}

func (c *optionCharset) apply(options *options) {
	options.charset = c.charset
}
