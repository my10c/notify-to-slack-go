// BSD 3-Clause License
//
// Copyright (c) 2023, © Badassops LLC / Luc Suryo
// All rights reserved.
//
// Version	:	0.1
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
		UserEmoji string
		MsgEmoji  string
		Url       string
	}
)

var (
   MyVersion   = "0.0.2"
    now         = time.Now()
    MyProgname  = path.Base(os.Args[0])
    myAuthor    = "Luc Suryo"
    myCopyright = "Copyright 2023 - " + strconv.Itoa(now.Year()) + " ©Badassops LLC"
    myLicense   = "License 3-Clause BSD, https://opensource.org/licenses/BSD-3-Clause ♥"
    myEmail     = "<luc@badassops.com>"
    MyInfo      = fmt.Sprintf("%s (version %s)\n%s\n%s\nWritten by %s %s\n",
        MyProgname, MyVersion, myCopyright, myLicense, myAuthor, myEmail)
    MyDescription = "Simple script send a message to a slack channel via a piped message."

	SlackConfigFile = "/etc/slack.conf"
	EmojiDefault    = ":badass:"
 	MsgEmojiDefault = ":red-alert:"
)
