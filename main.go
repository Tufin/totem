package main

import (
	"os"

	"github.com/tufin/logrus"
	"github.com/tufin/totem/analysis"
	"github.com/tufin/totem/common"
)

func main() {

	filePathToInvalidImports := analysis.NewCrawler(common.GetEnvOrExit("PACKAGE")).Run(common.GetEnvOrExit("PATH"))
	for currFilePath, currInvalidImports := range filePathToInvalidImports {
		logrus.Printf("'%s': '%v'", currFilePath, currInvalidImports)
	}

	if len(filePathToInvalidImports) > 0 {
		os.Exit(1)
	}
}
