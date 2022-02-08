package guide

import (
	`github.com/storezhang/gex`
)

var _ = envsParse

func envsParse() (int, error) {
	return gex.Exec(`redis`, gex.Envs(gex.ParseEnvs(`USERNAME=storezhang`)...))
}
