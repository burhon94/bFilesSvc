package servces

import "os"

func EnvOrFlag(envName string, flag *string) (value string, ok bool) {
	if flag == nil {
		return *flag, true
	}

	return os.LookupEnv(envName)
}