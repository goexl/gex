package gex

import (
	`strings`
)

var (
	_ = NewEnv
	_ = ParseEnv
)

type env struct {
	key   string
	value string
}

// NewEnv 创建环境变量
func NewEnv(key string, value string) *env {
	return &env{
		key:   key,
		value: value,
	}
}

// ParseEnv 解析环境变量
func ParseEnv(from string) (_env *env) {
	data := strings.Split(from, `=`)
	if 2 != len(data) {
		return
	}
	_env = &env{
		key:   data[0],
		value: data[1],
	}

	return
}
