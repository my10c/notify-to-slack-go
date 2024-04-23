module main

go 1.22.2

require (
	configurator v0.0.0
	getargs v0.0.0
	help v0.0.0 // indirect
	message v0.0.0
	vars v0.0.0
)

require github.com/my10c/packages-go/print v0.0.0-20230717011209-51a83962742b

require (
	github.com/BurntSushi/toml v1.1.0
	github.com/akamensky/argparse v1.4.0
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/slack-go/slack v0.12.2
)

replace configurator => ./mod/configurator

replace getargs => ./mod/getargs

replace help => ./mod/help

replace message => ./mod/message

replace vars => ./mod/vars
