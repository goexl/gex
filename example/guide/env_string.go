package guide

import (
	`github.com/goexl/gex`
)

var (
	_ = envString
	_ = envsString
)

func envString() (int, error) {
	return gex.Exec(`redis`, gex.StringEnv(`USERNAME=storezhang`))
}

func envsString() (int, error) {
	return gex.Exec(`redis`, gex.StringEnvs(`USERNAME=storezhang`, `PASSWORD=storezhang`))
}
