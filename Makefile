
CODE_NAME = notify-to-slack
SOURCES = $(CODE_NAME).go
BUILT_SOURCES = $(SOURCES)

all: clean build

build:	notify-to-slack.go
	go build -ldflags "-w -s" -o $(CODE_NAME) $(CODE_NAME).go

clean:
	@rm -f notify-to-slack
