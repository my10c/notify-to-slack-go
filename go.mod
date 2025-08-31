module main

go 1.24.4

require (
	configurator v0.0.0
	getargs v0.0.0
	help v0.0.0 // indirect
	logs v0.0.0 // indirect
	message v0.0.0
	vars v0.0.0
)

require (
	github.com/BurntSushi/toml v1.1.0
	github.com/akamensky/argparse v1.4.0
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/my10c/packages-go/print v0.0.0-20230717011209-51a83962742b
	github.com/slack-go/slack v0.12.2
	github.com/mitchellh/go-ps v1.0.0 // indirect
	github.com/my10c/packages-go/is v0.0.0-20230717011209-51a83962742b // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)


replace configurator => ./mod/configurator

replace getargs => ./mod/getargs

replace help => ./mod/help

replace logs => ./mod/logs

replace message => ./mod/message

replace vars => ./mod/vars
