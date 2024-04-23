// BSD 3-Clause License
//
// Copyright (c) 2023, © Badassops LLC / Luc Suryo
// All rights reserved.
//
// Version	:	0.1
//

package message

import (
	"fmt"
	"os"
	"io"
	"strings"

	// local
	"vars"

	// on github
	"github.com/slack-go/slack"
)

func SendMessage(msg string, config vars.SlackConfig) bool {
 	// create a new connection
 	slackAPI := slack.New(config.Token)

 	// setup the message options
 	slackMsgOptions := slack.PostMessageParameters{
 	 	Username:       config.User,
         IconEmoji:     config.UserEmoji,
         Markdown:      true,
         EscapeText:    false,
 	}

 	// send the message
 	_, _, err := slackAPI.PostMessage(config.Channel,
 					slack.MsgOptionText(msg, false),
 					slack.MsgOptionPostMessageParameters(slackMsgOptions),)
 	if err !=nil {
 		fmt.Printf("\nError sending the message, error %s..\n", err.Error())
		return false
 	}
	return true
}

func GetMessage(config vars.SlackConfig) string {
	//
	// 🟩 📡 🔴 🟢 
	//
	//	# HOSTSTATE				UP DOWN UNREACHABLE
	//	# SERVICESTATE			OK WARNING UNKNOWN CRITICAL <-- no longer needed, SERVICEOUTPUT have the info needed
	//	# SERVICEOUTPUT			[OK WARNING UNKNOWN CRITICAL] some text
	//	# SERVICEDISPLAYNAME	Alias of the check
	//
	var message string
	var notification_type string
	var notification_host string
	var notification_state string
	var service_name string

	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	data := strings.Split(string(stdin), "\n")
	notification_type = strings.Split(data[0], " ")[0]
	notification_host = strings.Split(data[0], " ")[1]
	url := fmt.Sprintf("<%s|%s>", config.Url, notification_host, notification_host)
	switch notification_type {
		case "Host:":
			// build the message for a host notification
			notification_state = strings.ReplaceAll(data[2], "HostState: ", "")
			if strings.Contains(data[1],"DOWN") {
				message = fmt.Sprintf(":red_circle: %s\n\n ⦿ host *alert* \n ⦿ DOWN\n ⦿ %s",
				 url, notification_state)
			}
			if strings.Contains(data[1], "UP") {
				message = fmt.Sprintf(":large_green_circle: %s\n\n ⦿ host *recovered* \n ⦿ UP\n ⦿ %s",
				 url, notification_state)
			}
		case "ServiceHost:":
			notification_state = strings.ReplaceAll(data[1], "ServiceOutput: ", "")
			service_name = strings.ReplaceAll(data[2], "ServiceName: ", "")
			if strings.Contains(data[1],"OK") {
				message = fmt.Sprintf(":large_green_circle: %s\n\n ⦿ Service *recovered* \n ⦿ %s\n ⦿ %s",
				 url, service_name, notification_state)
			} else {
				message = fmt.Sprintf(":red_circle: %s\n\n ⦿ Service *alert* \n ⦿ %s\n ⦿ %s",
				 url, service_name, notification_state)
			}
		default:
			message = ":red_alert: unknown error occured, please check the monitor dashboard"
	}
	return message
}
