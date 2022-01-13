package gex

import (
	`os`
)

type (
	fileOption interface {
		applyFile(options *options)
	}

	fileOptions struct {
		mode os.FileMode
	}
)

func defaultFileOptions() *fileOptions {
	return &fileOptions{
		mode: os.ModePerm,
	}
}
