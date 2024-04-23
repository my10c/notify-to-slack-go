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
	"time"
	"regexp"

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
	var msg string
	result := make(chan string, 1)
	go func() {
		result <- getMessage(config)
	} ()
	select {
		// we should get data within 2 seconds
		// otherwise we exit
		case <-time.After(2 * time.Second):
			os.Exit(3)
		case msg = <-result:
			break
	}
	return msg
}

func getMessage(config vars.SlackConfig) string {
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
	// we know we need at least 60 chars, exit if we get less then 50 chars
	if len(string(stdin)) < 50 {
		os.Exit(3)
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
			re := regexp.MustCompile("status.*")
			url := fmt.Sprintf("<%s|monitor home>", re.ReplaceAllString(config.Url, "main.cgi"))
			message = fmt.Sprintf(":large_yellow_circle: %s\n\n ⦿ unknown *error* occured \n ⦿ please check the monitor dashboard\n",url)
	}
	return message
}
