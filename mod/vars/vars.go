// BSD 3-Clause License
//
// Copyright (c) 2023 - 2025, © Badassops LLC / Luc Suryo
// All rights reserved.
//
// Version	:	0.2
//

package vars

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"time"
)

type (
	SlackConfig struct {
		Token     string
		User      string
		Channel   string
		UserEmoji string `default: ":heart-monitor:"`
		MsgEmoji  string `default: ":alert:"`
		Url       string
	}
	LogConfig struct {
		LogEnable     bool   `default: false`
		LogsDir       string `default: "fmt.Sprintf("/var/tmp/%s", MyProgname)"`
		LogFile       string `default: "fmt.Sprintf("%s.log", MyProgname)"`
		LogMaxSize    int    `default: 32`
		LogMaxBackups int    `default: 7`
		LogMaxAge     int    `default: 7`
	}
	Auth struct {
		AllowUsers []string
		AllowMods  []string
	}
	GivenArgs struct {
		ConfigFile   string
		SlackMessage []string
		Quite        bool
		DebugMode    bool
	}
)

var (
	Off    = "\x1b[0m"    // Text Reset
	Black  = "\x1b[1;30m" // Black
	Red    = "\x1b[1;31m" // Red
	Green  = "\x1b[1;32m" // Green
	Yellow = "\x1b[1;33m" // Yellow
	Blue   = "\x1b[1;34m" // Blue
	Purple = "\x1b[1;35m" // Purple
	Cyan   = "\x1b[1;36m" // Cyan
	White  = "\x1b[1;37m" // White

	RedUnderline = "\x1b[4;31m" // Red underline
	OneLineUP    = "\x1b[A"
)

var (
	MyVersion   = "0.2.0"
	now         = time.Now()
	MyProgname  = path.Base(os.Args[0])
	myAuthor    = "Luc Suryo"
	myCopyright = "Copyright 2023 - " + strconv.Itoa(now.Year()) + " ©Badassops LLC"
	myLicense   = "License 3-Clause BSD, https://opensource.org/licenses/BSD-3-Clause ♥"
	myEmail     = "<luc@badassops.com>"
	MyInfo      = fmt.Sprintf("%s (version %s)\n%s\n%s\nWritten by %s %s\n",
		MyProgname, MyVersion, myCopyright, myLicense, myAuthor, myEmail)
	MyDescription = "Simple script send a message to a slack channel via a piped message or by the given the message on the command line."

	// configuration file and slack default values
	DefaultQuiet         = false
	ConfigFile           = "/usr/local/etc/slack/slack.conf"
	DefaultUserEmoji     = ":heart-monitor:"
	DefaultMsgEmoji      = ":alert:"
	DefaultLogEnable     = false
	DefaultLogsDir       = fmt.Sprintf("/var/tmp/%s", MyProgname)
	DefaultLogFile       = fmt.Sprintf("%s.log", MyProgname)
	DefaultLogMaxSize    = 32
	DefaultLogMaxBackups = 7
	DefaultLogMaxAge     = 7
)
