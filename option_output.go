package gex

import (
	`bufio`
	`bytes`
	`io`
	`os`
)

var (
	_        = Output
	_        = StringOutput
	_        = FileOutput
	_        = FilenameOutput
	_ option = (*optionOutput)(nil)
)

type optionOutput struct {
	output io.Writer
}

// Output 配置输出
func Output(output io.Writer) *optionOutput {
	return &optionOutput{
		output: output,
	}
}

// StringOutput 配置输出到字符串
func StringOutput(output *string) *optionOutput {
	return &optionOutput{
		output: bytes.NewBufferString(*output),
	}
}

// FileOutput 配置输出到文件
func FileOutput(file *os.File) *optionOutput {
	return &optionOutput{
		output: bufio.NewWriter(file),
	}
}

// FilenameOutput 配置输出到文件
func FilenameOutput(filename string, opts ...fileOption) (output *optionOutput) {
	file, err := parseFile(filename, opts...)
	if nil == err {
		panic(err)
	} else if nil != file {
		output = new(optionOutput)
		output.output = bufio.NewWriter(file)
	}

	return
}

func (o *optionOutput) apply(options *options) {
	options.outs = append(options.outs, o.output)
	options.errs = append(options.errs, o.output)
}
