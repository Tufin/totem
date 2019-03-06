package main

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/tufin/totem/analysis"
	"github.com/tufin/totem/common"
)

func init() {

	logrus.SetReportCaller(true)
	logrus.SetLevel(getLogLevel())
}

func main() {

	filePathToInvalidImports := getInvalidImports()

	// print
	for currFilePath, currInvalidImports := range filePathToInvalidImports {
		logrus.Printf("'%s': '%v'", currFilePath, currInvalidImports)
	}

	// exit
	if len(filePathToInvalidImports) > 0 {
		os.Exit(1)
	}
}

func getInvalidImports() map[string][]string {

	imports := common.NewList()
	imports.AddItems(strings.Split(common.GetEnv("TOTEM_COMMON_IMPORTS"), ","))

	crawler := analysis.NewCrawler(common.GetEnvOrExit("TOTEM_PACKAGE"), imports)
	root := common.GetEnvOrExit("TOTEM_PATH")
	var ret map[string][]string
	if service := common.GetEnv("TOTEM_SERVICE"); service != "" {
		ret = crawler.RunService(root, service)
	} else {
		ret = crawler.Run(root)
	}

	return ret
}

func getLogLevel() logrus.Level {

	ret := logrus.InfoLevel
	level := common.GetEnv("LOG_LEVEL")
	if level != "" {
		if strings.EqualFold(level, "debug") {
			ret = logrus.DebugLevel
		} else if strings.EqualFold(level, "warn") {
			ret = logrus.WarnLevel
		} else if strings.EqualFold(level, "error") {
			ret = logrus.ErrorLevel
		}
	}

	return ret
}
