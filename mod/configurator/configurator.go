// BSD 3-Clause License
//
// Copyright (c) 2023, © Badassops LLC / Luc Suryo
// All rights reserved.
//
// Version	:	0.1
//

package configurator

import (
	"fmt"
	"os"

	// local
    "vars"

	// on github	
	"github.com/my10c/packages-go/print"
	"github.com/BurntSushi/toml"
)

type (
	tomlConfig struct {
		Slack vars.SlackConfig `toml:"slack"`
	}
)

var (
	Print = print.New()
)

func GetConfig() vars.SlackConfig {
	var configValues tomlConfig
	var configured vars.SlackConfig

	if _, err := toml.DecodeFile(vars.SlackConfigFile, &configValues); err != nil {
		Print.PrintRed("Error reading the configuration file\n")
		fmt.Fprintln(os.Stderr, err)
        Print.PrintBlue("Aborting...\n")
		os.Exit(1)
	}
	if	len(configValues.Slack.Token)   == 0 ||
		len(configValues.Slack.User)    == 0 ||
		len(configValues.Slack.Channel) == 0 ||
		len(configValues.Slack.Url)     == 0 {
		Print.PrintRed("Error reading the configuration file, some value are missing or is empty\n")
		Print.PrintBlue("Make sure token, user, channel and url are set\n")
        Print.PrintBlue("Aborting...\n")
        os.Exit(1)
	}
	configured.Token   = configValues.Slack.Token
	configured.User    = configValues.Slack.User
	configured.Channel = configValues.Slack.Channel
	configured.Url     = configValues.Slack.Url
	// We set to default value and later change it is it was set 
	// in the configuration file
	configured.UserEmoji = vars.EmojiDefault
	configured.MsgEmoji	 = vars.MsgEmojiDefault
	if len(configValues.Slack.UserEmoji) !=0 {
		configured.UserEmoji = configValues.Slack.UserEmoji
	}
	if len(configValues.Slack.MsgEmoji) !=0 {
		configured.MsgEmoji = configValues.Slack.MsgEmoji
	}
	return configured	
}
