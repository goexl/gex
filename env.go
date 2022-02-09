package gex

import (
	`strings`
)

var (
	_ = NewEnv
	_ = ParseEnvs
)

type env struct {
	key   string
	value string
}

// NewEnv 创建环境变量
func NewEnv(key string, value interface{}) *env {
	return &env{
		key:   key,
		value: toString(value)[0],
	}
}

// ParseEnvs 解析环境变量
func ParseEnvs(from ...string) (envs []*env) {
	envs = make([]*env, 0, len(from))
	for _, _env := range from {
		data := strings.Split(_env, `=`)
		if 2 != len(data) {
			continue
		}

		envs = append(envs, &env{
			key:   data[0],
			value: data[1],
		})
	}

	return
}
