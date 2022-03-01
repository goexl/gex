package gex

import (
	`os`
)

type (
	fileOption interface {
		applyFile(options *fileOptions)
	}

	fileOptions struct {
		mode      os.FileMode
		writeMode writeMode
	}
)

func defaultFileOptions() *fileOptions {
	return &fileOptions{
		mode: os.ModePerm,
	}
}
