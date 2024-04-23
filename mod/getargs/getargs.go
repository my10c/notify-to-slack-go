// BSD 3-Clause License
//
// Copyright (c) 2023, © Badassops LLC / Luc Suryo
// All rights reserved.
//
// Version	:	0.1
//

package getargs

import (
	"fmt"
	"os"

	// local
	"help"
    "vars"

	// on github	
	"github.com/akamensky/argparse"
)

func GetArgs() bool {
	parser := argparse.NewParser(vars.MyProgname, vars.MyDescription)
	showVersion := parser.Flag("v", "version",
		&argparse.Options{
			Required: false,
			Help:	"Show version",
		})
	showInfo := parser.Flag("i", "info",
		&argparse.Options{
			Required: false,
			Help:     "Show how to use " + vars.MyProgname,
		})
	debug := parser.Flag("t", "test",
		&argparse.Options{
			Required:   false,
			Help:   	"test mode, no message will be sent",
			Default:	false,
		})
	showSetup := parser.Flag("s", "setup",
		&argparse.Options{
			Required: false,
			Help:     "Show how to setup in nagios or naemon",
		})
	showSetupConfig := parser.Flag("S", "slack-config",
		&argparse.Options{
			Required: false,
			Help:     "Show how to setup the slack configuration file",
		})

	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Printf(parser.Usage(err))
		os.Exit(1)
	}
	if *showVersion {
		help.Version()
	}
	if *showInfo {
		help.Info()
	}
	if *showSetup {
		help.Setup()
	}
	if *showSetupConfig {
		help.SetupConfig()
	}
	if *debug {
		return true
    }
	return false
}
