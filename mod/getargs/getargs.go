// BSD 3-Clause License
//
// Copyright (c) 2023 - 2025, © Badassops LLC / Luc Suryo
// All rights reserved.
//

package getargs

import (
	"os"

	// local
	"help"
	"vars"

	// on github
	"github.com/akamensky/argparse"
	"github.com/my10c/packages-go/is"
	"github.com/my10c/packages-go/print"
)

var (
	Is    = is.New()
	Print = print.New()
)

// func GetArgs() (bool, string, []string, bool) {
func GetArgs() *vars.GivenArgs {
	Args := &vars.GivenArgs{}
	parser := argparse.NewParser(vars.MyProgname, vars.MyDescription)
	configFile := parser.String("c", "configFile",
		&argparse.Options{
			Required: false,
			Help:     "Configuration file to be use",
			Default:  vars.ConfigFile,
		})
	showVersion := parser.Flag("v", "version",
		&argparse.Options{
			Required: false,
			Help:     "Show version",
		})
	showInfo := parser.Flag("i", "info",
		&argparse.Options{
			Required: false,
			Help:     "Show how to use " + vars.MyProgname,
		})
	debug := parser.Flag("t", "test",
		&argparse.Options{
			Required: false,
			Help:     "test mode, no message will be sent",
			Default:  false,
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
	slackMessage := parser.StringList("m", "message",
		&argparse.Options{
			Required: false,
			Help:     "Message to be sent between double quotes or single quotes, implies no stdin reading",
		})
	quietFlag := parser.Flag("q", "quiet",
		&argparse.Options{
			Required: false,
			Help:     "Quiet mode",
			Default:  vars.DefaultQuiet,
		})

	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		msg := parser.Usage(err)
		Print.PrintRed(msg)
		os.Exit(1)
	}

	if _, ok, _ := Is.IsExist(*configFile, "file"); !ok {
		msg := "Configuration file " + *configFile + " does not exist\n"
		Print.PrintRed(msg)
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

	Args.ConfigFile = *configFile
	Args.SlackMessage = *slackMessage
	Args.Quite = *quietFlag
	Args.DebugMode = *debug
	return Args
	// return *debug, *configFile, *slackMessage, *quietFlag
}
