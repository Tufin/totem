package main

import (
	"os"

	"github.com/tufin/logrus"
	"github.com/tufin/totem/analysis"
)

func main() {

	filePathToInvalidImports := analysis.Run("/Users/effi.bar/view/go/src/github.com/tufin/orca")
	for currFilePath, currInvalidImports := range filePathToInvalidImports {
		logrus.Printf("'%s': '%v'", currFilePath, currInvalidImports)
	}

	if len(filePathToInvalidImports) > 0 {
		os.Exit(1)
	}

	//app := cli.NewApp()
	//app.Name = "totem"
	//app.Usage = "find invalid imports"
	//app.Version = "0.1"
	//
	//app.Commands = []cli.Command{
	//	{
	//		Name:    "find",
	//		Aliases: []string{"f"},
	//		Usage:   "find invalid imports",
	//		Subcommands: cli.Commands{
	//			cli.Command{
	//				Name: "token",
	//				Flags: []cli.Flag{cli.StringFlag{Name: "scope, s"},
	//					cli.StringFlag{Name: "label, l"},
	//					cli.StringFlag{Name: "domain, d", EnvVar: "DOMAIN"},
	//					cli.StringFlag{Name: "project, p", EnvVar: "PROJECT"},
	//				},
	//			},
	//		},
	//	},
	//}
	//
	//app.Run()
}
