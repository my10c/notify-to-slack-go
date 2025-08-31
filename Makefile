
CODE_NAME = notify-to-slack
SOURCES = $(CODE_NAME).go \
	mod/configurator/configurator.go \
	mod/getargs/getargs.go \
	mod/help/help.go \
	mod/logs/logs.go \
	mod/message/message.go \
	mod/vars/vars.go \

BUILT_SOURCES = $(SOURCES)

all: clean build

build:	notify-to-slack.go
	go build -ldflags "-w -s" -o $(CODE_NAME) $(CODE_NAME).go

clean:
	@rm -f notify-to-slack
