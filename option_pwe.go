package gex

var (
	_        = Pwe
	_        = DisablePwe
	_ option = (*optionPwe)(nil)
)

type optionPwe struct {
	pwe bool
}

// Pwe 当出错时打印信息到控制台
func Pwe() *optionPwe {
	return &optionPwe{
		pwe: true,
	}
}

// DisablePwe 当出错时静默处理
func DisablePwe() *optionPwe {
	return &optionPwe{
		pwe: false,
	}
}

func (p *optionPwe) apply(options *options) (err error) {
	options.pwe = p.pwe

	return
}
