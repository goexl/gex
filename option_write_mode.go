package gex

var (
	_            = WriteMode
	_ fileOption = (*optionWriteMode)(nil)
)

type optionWriteMode struct {
	mode writeMode
}

// WriteMode 配置复制模式
func WriteMode(mode writeMode) *optionWriteMode {
	return &optionWriteMode{
		mode: mode,
	}
}

func (wm *optionWriteMode) applyFile(options *fileOptions) {
	options.writeMode = wm.mode
}
