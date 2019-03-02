package main

import "github.com/tufin/totem/analysis"

func main() {

	analysis.Run("/Users/effi.bar/view/go/src/github.com/tufin/orca")
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
