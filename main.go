package main

import (
	"os"
	"strings"

	"github.com/tufin/logrus"
	"github.com/tufin/totem/analysis"
	"github.com/tufin/totem/common"
)

func main() {

	imports := common.NewList()
	imports.AddItems(strings.Split(common.GetEnv("COMMON_IMPORTS"), ","))

	filePathToInvalidImports := analysis.NewCrawler(
		common.GetEnvOrExit("PACKAGE"), imports).
		Run(common.GetEnvOrExit("PATH"))

	for currFilePath, currInvalidImports := range filePathToInvalidImports {
		logrus.Printf("'%s': '%v'", currFilePath, currInvalidImports)
	}

	if len(filePathToInvalidImports) > 0 {
		os.Exit(1)
	}
}
