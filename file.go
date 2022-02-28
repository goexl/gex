package exec

import (
	`fmt`
	`os`
	`path/filepath`
	`strings`
)

func parseFile(filename string, opts ...fileOption) (file *os.File, err error) {
	_options := defaultFileOptions()
	for _, opt := range opts {
		opt.applyFile(_options)
	}

	if exist(filename) {
		switch _options.writeMode {
		case WriteModeError:
			err = errFileExist
		case WriteModeSkip:
			return
		case WriteModeOverride:
			file, err = os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, _options.mode)
		case WriteModeRename:
			file, err = os.OpenFile(newFilename(filename), os.O_WRONLY|os.O_CREATE, _options.mode)
		case WriteModeAppend:
			file, err = os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, _options.mode)
		}
	}

	return
}

func exist(filename string) (exist bool) {
	if _, err := os.Stat(filename); nil != err && os.IsNotExist(err) {
		exist = false
	} else {
		exist = true
	}

	return
}

func newFilename(path string) (filename string) {
	for {
		index := 1
		filename = filepath.Join(filepath.Dir(path), name(path, fmt.Sprintf(`%d.%s`, index, filepath.Ext(path))))
		if !exist(filename) {
			break
		}
	}

	return
}

func name(path string, ext string) (filename string) {
	base := filepath.Base(path)
	_ext := filepath.Ext(path)
	_name := base[:len(base)-len(_ext)]

	if `` == strings.TrimSpace(ext) {
		filename = _name
	} else {
		filename = fmt.Sprintf(`%s.%s`, _name, ext)
	}

	return
}
