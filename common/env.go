package common

import (
	"os"

	"github.com/tufin/logrus"
)

func GetEnvOrExit(key string) string {

	ret := os.Getenv(key)
	if ret == "" {
		logrus.Fatalf("Please, set '%s'", key)
	}
	logrus.Infof("'%s': '%s'", key, ret)

	return ret
}

func GetEnv(key string) string {

	ret := os.Getenv(key)
	logrus.Infof("'%s': '%s'", key, ret)

	return ret
}
