// BSD 3-Clause License
//
// Copyright (c) 2023, © Badassops LLC / Luc Suryo
// All rights reserved.
//
// Version	:	0.1
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
	Print.PrintYellow(vars.MyProgname + " usage should not be use with any flags, unless you want to:\n")
	Print.PrintYellow("see this information 😈 (-i), test without actually sent the message 🤣 (-t),\n")
	Print.PrintYellow("- show to configure Naemon/Nagios (-s)\n")
	Print.PrintYellow("- show to configure slack configuratiom file (" + vars.SlackConfigFile + ") (-S)\n")
	Print.PrintYellow("- see the version (-v)\n")
	Print.PrintPurple("It should be use with pipped data from a nagios or naemon command.\n")
	Print.PrintPurple("Example: /usr/bin/printf \"%s\" \"<some-data>\" | " +  vars.MyProgname + "\n\n")
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
	Print.PrintYellow("The configuration file is: " + vars.SlackConfigFile + "\n")
	Print.PrintPurple("\n[slack]\n")
	Print.PrintPurple("# these are required\n")
	Print.PrintPurple("token       = \"xoxb-xxx-xxx-xxx\"\n")
	Print.PrintPurple("user        = \"some-bot-id\"\n")
	Print.PrintPurple("channel     = \"some-slack-channel\"\n")
	Print.PrintPurple("url         = \"url-of-your-nagios/naemon\"\n")
	Print.PrintGreen("# these are optional, the default is shown below\n")
	Print.PrintGreen("# emoji of the user\n")
	Print.PrintGreen("userEmoji   = \":badass:\"\n")
	Print.PrintGreen("# message emoji\n")
	Print.PrintGreen("msgEmoji    = \":red-alert:\"\n\n")
	Print.PrintBlue("url example: \"https://naemon.your-domain.tld/thruk/cgi-bin/status.cgi?host=\"\n\n")
	os.Exit(0)
}
