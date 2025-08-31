// BSD 3-Clause License
//
// Copyright (c) 2023 - 2025, © Badassops LLC / Luc Suryo
// All rights reserved.
//

package message

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"time"

	// local
	logs "logs"
	"vars"

	// on github
	"github.com/slack-go/slack"
)

func SendMessage(msg string, config vars.SlackConfig) bool {
	// create a new connection
	slackAPI := slack.New(config.Token)

	// setup the message options
	slackMsgOptions := slack.PostMessageParameters{
		Username:   config.User,
		IconEmoji:  config.UserEmoji,
		Markdown:   true,
		EscapeText: false,
	}

	// send the message
	_, _, err := slackAPI.PostMessage(config.Channel,
		slack.MsgOptionText(msg, false),
		slack.MsgOptionPostMessageParameters(slackMsgOptions))
	if err != nil {
		msg := fmt.Sprintf("\nError sending the message, error %s..\n", err.Error())
		fmt.Printf(msg)
		logs.LogIt(msg, "ERROR", true)
		return false
	}
	return true
}

func GetMessage(config vars.SlackConfig) string {
	var msg string
	result := make(chan string, 1)
	go func() {
		result <- getMessage(config)
	}()
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
	// 🟩 🔴 🟢
	//
	//	# HOSTSTATE				UP DOWN UNREACHABLE
	//	# SERVICESTATE			OK WARNING UNKNOWN CRITICAL <-- no longer needed, SERVICEOUTPUT have the info needed
	//	# SERVICEOUTPUT			[OK WARNING UNKNOWN CRITICAL] some text
	//	# SERVICEDISPLAYNAME	Alias of the check
	//
	var message string
	var service_name string

	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	// we know we need at least 60 chars, exit if we get less then 50 chars
	// if get nothing we need to exist
	if len(string(stdin)) < 50 {
		fmt.Printf("🌕 %s\n\n ⦿ unknown *error* occured \n ⦿ please check the monitor dashboard\n",
			config.Url)
		os.Exit(3)
	}
	data := strings.Split(string(stdin), "\n")
	notification_type := strings.Split(data[0], " ")[0]
	notification_host := strings.Split(data[0], " ")[1]
	notification_msg := "no message provided"
	if len(data[3]) > 1 {
		notification_msg = strings.ReplaceAll(data[3], "ServiceState: ", "")
	}
	url := fmt.Sprintf("Monitorl URL: <%s|%s>", config.Url+notification_host, notification_host)
	switch notification_type {
	case "Host:":
		// build the message for a host notification
		notification_state := strings.ReplaceAll(data[2], "HostState: ", "")
		if strings.Contains(data[1], "DOWN") {
			message = fmt.Sprintf(":alert: %s\n\n ⦿ host *alert* \n ⦿ DOWN\n ⦿ %s\n⦿ %s",
				url, notification_state, notification_msg)
		}
		if strings.Contains(data[1], "UP") {
			message = fmt.Sprintf(":goodgreen: %s\n\n ⦿ host *recovered* \n ⦿ UP\n ⦿ %s\n⦿ %s",
				url, notification_state, notification_msg)
		}
	case "ServiceHost:":
		// build the message for a service notification
		notification_state := strings.ReplaceAll(data[1], "ServiceOutput: ", "")
		service_name = strings.ReplaceAll(data[2], "ServiceName: ", "")
		if strings.Contains(data[1], "OK") {
			message = fmt.Sprintf(":goodgreen: %s\n\n ⦿ Service *recovered* \n ⦿ %s\n ⦿ %s\n⦿ %s",
				url, service_name, notification_state, notification_msg)
		} else {
			message = fmt.Sprintf(":alert: %s\n\n ⦿ Service *alert* \n ⦿ %s\n ⦿ %s\n⦿ %s",
				url, service_name, notification_state, notification_msg)
		}
	default:
		// build the message for a unknown message
		re := regexp.MustCompile("status.*")
		url := fmt.Sprintf("<%s|monitor home>", re.ReplaceAllString(config.Url, "main.cgi"))
		message = fmt.Sprintf(":large_yellow_circle: %s\n\n ⦿ unknown *error* occured \n ⦿ please check the monitor dashboard\n", url)
	}
	return message
}
