package analysis

import (
	"os"

	"github.com/tufin/logrus"
)

func GetEnv(key string) string {

	ret := os.Getenv(key)
	logrus.Infof("'%s': '%s'", key, ret)

	return ret
}
