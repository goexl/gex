package pipe

import (
	`github.com/storezhang/gex`
)

var _ = envsParse

func envsParse() (int, error) {
	return gex.Exec(`redis`, gex.Pipe(
		`echo`,
		gex.Args(`www.163.com`),
		gex.Envs(gex.ParseEnvs(`USERNAME=storezhang`)...),
	))
}
