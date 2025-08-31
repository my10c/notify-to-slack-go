// BSD 3-Clause License
//
// Copyright (c) 2023 - 2025, © Badassops LLC / Luc Suryo
// All rights reserved.
//

package help

import (
	"fmt"
	"os"

	// local
	"vars"

	// on github
	"github.com/my10c/packages-go/print"
)

var (
	Print = print.New()
)

func Version() {
	Print.ClearScreen()
	Print.PrintCyan("Version: " + vars.MyVersion + "\n")
	Print.PrintPurple(vars.MyInfo + "\n")
	os.Exit(0)
}

func Info() {
	Print.ClearScreen()
	Print.PrintCyan(vars.MyProgname + " usage should not be use with any flags, unless you want to:\n")
	Print.PrintCyan("see this information 😈 (-i), test without actually sent the message 🤣 (-t),\n")
	Print.PrintCyan("- show to configure Naemon/Nagios (-s)\n")
	Print.PrintCyan("- show to configure slack configuratiom file (" + vars.ConfigFile + ") (-S)\n")
	Print.PrintCyan("- see the version (-v)\n")
	Print.PrintPurple("It should be use with pipped data from a nagios or naemon command.\n")
	Print.PrintPurple("Example: /usr/bin/printf \"%s\" \"<some-data>\" | " + vars.MyProgname + "\n\n")
	os.Exit(0)
}

func Setup() {
	Print.ClearScreen()
	Print.PrintGreen("**Do** note that the \\n are required! It is use to parse the message\n")
	Print.PrintBlue("The script depends of the variables passed and their order!\n")

	fmt.Printf("\n# notify-host-to-slack command definition\n")
	fmt.Printf("define command{\n")
	fmt.Printf("  command_name notify-host-to-slack\n")
	fmt.Printf("  command_line /usr/bin/printf \"%%b\"")
	fmt.Printf(" \"Host: $HOSTNAME$\\nHostOutput: $HOSTOUTPUT$\\nHostState: $HOSTSTATE$\\n\"")
	fmt.Printf(" | /usr/local/sbin/notify-to-slack 2>> /tmp/hosts_notification.log\n}\n\n")

	Print.PrintGreen("and\n")

	fmt.Printf("\n# notify-service-to-slack command definition\n")
	fmt.Printf("#define command{\n")
	fmt.Printf("  command_name notify-service-by-teams\n")
	fmt.Printf("  command_line  /usr/bin/printf \"%%b\"")
	fmt.Printf(" \"ServiceHost: $HOSTNAME$\\nServiceOutput: $SERVICEOUTPUT$\\n")
	fmt.Printf("ServiceName: $SERVICEDISPLAYNAME$\\nServiceState: $SERVICESTATE$\\n\"")
	fmt.Printf(" | /usr/local/sbin/notify-to-slack 2>> /tmp/services_notification.log\n}\n\n")

	os.Exit(0)
}

func SetupConfig() {
	Print.ClearScreen()
	Print.PrintCyan("The configuration file is: " + vars.ConfigFile + "\n")
	Print.PrintPurple("\n[slack]\n")
	Print.PrintPurple("# these are required\n")
	Print.PrintPurple("token       = \"xoxb-xxx-xxx-xxx\"\n")
	Print.PrintPurple("user        = \"some-bot-id\"\n")
	Print.PrintPurple("channel     = \"some-slack-channel\"\n")
	Print.PrintPurple("url         = \"url-of-your-nagios/naemon\"\n")
	Print.PrintGreen("\n# these are optional, the default is shown below\n")
	Print.PrintGreen("# emoji of the user\n")
	Print.PrintGreen(fmt.Sprintf("userEmoji   = %v\n", vars.DefaultUserEmoji))
	Print.PrintGreen("# message emoji\n")
	Print.PrintGreen(fmt.Sprintf("msgEmoji    = %v\n", vars.DefaultMsgEmoji))
	Print.PrintBlue("url example: \"https://naemon.your-domain.tld/thruk/cgi-bin/status.cgi?host=\"\n\n")
	Print.PrintCyan("\nBelow the log configurations\n")
	Print.PrintGreen("logs is disable by default\n")
	Print.PrintGreen("logMaxBackups is count and logMaxAge in days\n")
	Print.PrintPurple("[logconfig]\n")
	Print.PrintPurple(fmt.Sprintf("enableLog       = %v\n", vars.DefaultLogEnable))
	Print.PrintPurple(fmt.Sprintf("logsDir         = %v\n", vars.DefaultLogsDir))
	Print.PrintPurple(fmt.Sprintf("logFile         = %v\n", vars.DefaultLogFile))
	Print.PrintPurple(fmt.Sprintf("logMaxSize      = %v\n", vars.DefaultLogMaxSize))
	Print.PrintPurple(fmt.Sprintf("logMaxBackups   = %v\n", vars.DefaultLogMaxBackups))
	Print.PrintPurple(fmt.Sprintf("logMaxAge       = %v\n", vars.DefaultLogMaxAge))

	os.Exit(0)
}
