// BSD 3-Clause License
//
// Copyright (c) 2023 - 2025, © Badassops LLC / Luc Suryo
// All rights reserved.
//

package configurator

import (
	"fmt"
	"os"

	// local
	"vars"

	// on github
	"github.com/BurntSushi/toml"
	"github.com/my10c/packages-go/print"
)

type (
	tomlConfig struct {
		Auth      vars.Auth        `toml:"auth"`
		Slack     vars.SlackConfig `toml:"slack"`
		LogConfig vars.LogConfig   `toml:"logconfig"`
	}
)

var (
	Print = print.New()
)

// func GetConfig() vars.SlackConfig {
func GetConfig(configFile string) tomlConfig {
	var configValues tomlConfig

	if _, err := toml.DecodeFile(configFile, &configValues); err != nil {
		Print.PrintRed("Error reading the configuration file\n")
		fmt.Fprintln(os.Stderr, err)
		Print.PrintBlue("Aborting...\n")
		os.Exit(1)
	}
	// make sure all required configuration was set
	// slack
	if len(configValues.Slack.Token) == 0 ||
		len(configValues.Slack.User) == 0 ||
		len(configValues.Slack.Channel) == 0 ||
		len(configValues.Slack.Url) == 0 {
		Print.PrintRed("Error reading the configuration file, some value are missing or is empty\n")
		Print.PrintBlue("Make sure token, user, channel and url are set\n")
		Print.PrintBlue("Aborting...\n")
		os.Exit(1)
	}
	// auth == configuration file access
	if len(configValues.Auth.AllowUsers) == 0 ||
		len(configValues.Auth.AllowMods) == 0 {
		Print.PrintRed("Error reading the configuration file, some value are missing or is empty\n")
		Print.PrintBlue("Make sure allowUsers and allowMods are set\n")
		Print.PrintBlue("Aborting...\n")
		os.Exit(1)
	}
	// log is log configuration and set to default if not set
	if len(configValues.LogConfig.LogsDir) == 0 {
		configValues.LogConfig.LogsDir = vars.DefaultLogsDir
	}
	if len(configValues.LogConfig.LogFile) == 0 {
		configValues.LogConfig.LogFile = vars.DefaultLogFile
	}
	if configValues.LogConfig.LogMaxSize == 0 {
		configValues.LogConfig.LogMaxSize = vars.DefaultLogMaxSize
	}
	if configValues.LogConfig.LogMaxBackups == 0 {
		configValues.LogConfig.LogMaxBackups = vars.DefaultLogMaxBackups
	}
	if configValues.LogConfig.LogMaxAge == 0 {
		configValues.LogConfig.LogMaxAge = vars.DefaultLogMaxAge
	}
	return configValues
}
