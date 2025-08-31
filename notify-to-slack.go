// BSD 3-Clause License
//
// Copyright (c) 2023 - 2025, © Badassops LLC / Luc Suryo
// All rights reserved.
//
// Version	:	0.2
//

package main

import (
	"fmt"
	"os"
	"strings"

	// local
	conf "configurator"
	args "getargs"
	logs "logs"
	msg "message"
	vars "vars"

	// on github
	"github.com/my10c/packages-go/is"
	"github.com/my10c/packages-go/print"
)

func main() {
	isPtr := is.New()
	printPtr := print.New()

	// get given parameters
	//debug, configFile := args.GetArgs()
	args := args.GetArgs()

	// get the conguration in the config file
	config := conf.GetConfig(args.ConfigFile)

	// initialize the logger system is it was set to true
	if config.LogConfig.LogEnable {
		LogConfig := &vars.LogConfig{
			LogsDir:       config.LogConfig.LogsDir,
			LogFile:       config.LogConfig.LogFile,
			LogMaxSize:    config.LogConfig.LogMaxSize,
			LogMaxBackups: config.LogConfig.LogMaxBackups,
			LogMaxAge:     config.LogConfig.LogMaxAge,
		}
		logs.InitLogs(LogConfig)
	}

	// make sure the configuration file has the proper settings
	runningUser, _ := isPtr.IsRunningUser()
	if !isPtr.IsInList(config.Auth.AllowUsers, runningUser) {
		msg := fmt.Sprintf("The program has to be run as these user(s): %s or use sudo, aborting..\n",
			strings.Join(config.Auth.AllowUsers[:], ", "))
		printPtr.PrintRed(msg)
		logs.LogIt(msg, "ERROR", config.LogConfig.LogEnable)
		os.Exit(0)
	}
	ownerInfo, ownerOK := isPtr.IsFileOwner(args.ConfigFile, config.Auth.AllowUsers)
	if !ownerOK {
		msg := fmt.Sprintf("%s,\nAborting..\n", ownerInfo)
		printPtr.PrintRed(msg)
		logs.LogIt(msg, "ERROR", config.LogConfig.LogEnable)
		os.Exit(0)
	}
	permInfo, permOK := isPtr.IsFilePermission(args.ConfigFile, config.Auth.AllowMods)
	if !permOK {
		msg := fmt.Sprintf("%s,\nAborting..\n", permInfo)
		printPtr.PrintRed(msg)
		logs.LogIt(msg, "ERROR", config.LogConfig.LogEnable)
		os.Exit(0)
	}

	if len(args.SlackMessage) != 0 {
		// direct message
		slackMsg := strings.Join(args.SlackMessage, "")
		if msg.SendMessage(slackMsg, config.Slack) {
			if !args.Quite {
				fmt.Printf("Message: %v : has been sent\n", slackMsg)
			}
		} else {
			fmt.Printf("Message: %v : failed to been sent, check logs\n", slackMsg)
		}
	} else {
		// piped message
		slackMsg := msg.GetMessage(config.Slack)
		if args.DebugMode {
			fmt.Printf("%s\n", slackMsg)
		} else {
			msg.SendMessage(slackMsg, config.Slack)
		}
	}
	os.Exit(0)
}
