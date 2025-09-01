
MACHINE = generic

UNAME_S := $(shell uname -s)
UNAME_M := $(shell uname -m)

ifeq ($(UNAME_S),Linux)
    OS_ID = Linux_$(UNAME_M)
    GID = god
endif
ifeq ($(UNAME_S),Darwin)
    OS_ID = Darwin_$(UNAME_M)
    GID = staff
endif

CODE_NAME = notify-to-slack
SOURCES = $(CODE_NAME).go \
	mod/configurator/configurator.go \
	mod/getargs/getargs.go \
	mod/help/help.go \
	mod/logs/logs.go \
	mod/message/message.go \
	mod/vars/vars.go \

BUILT_SOURCES = $(SOURCES)
TOOL_VERSION := $(shell cat mod/vars/vars.go | grep MyVersion | egrep -v MyProgname | awk '{print $$3}')

all:	release/$(CODE_NAME)_$(OS_ID) \
		release/$(CODE_NAME)_$(OS_ID).tar.gz \
		release/$(CODE_NAME)_$(OS_ID).sha256

release/$(CODE_NAME)_$(OS_ID): $(BUILT_SOURCES)
	@echo "build the $(CODE_NAME)_$(OS_ID) binary..."
	@go build -o release/$(CODE_NAME)_$(OS_ID) $(CODE_NAME).go
	@echo "set owner and strip the binary"
	@chown luc:$(GID) release/$(CODE_NAME)_$(OS_ID)
	@strip release/$(CODE_NAME)_$(OS_ID)

release/$(CODE_NAME)_$(OS_ID).tar.gz: release/$(CODE_NAME)_$(OS_ID)
	@echo "create the $(CODE_NAME)_$(OS_ID).tar.gz archive..."
	@(cd release ; tar zcf $(CODE_NAME)_$(OS_ID).tar.gz $(CODE_NAME)_$(OS_ID))

release/$(CODE_NAME)_$(OS_ID).sha256: release/$(CODE_NAME)_$(OS_ID).tar.gz
	@echo "create the sha256 information file..."
	@sha256sum release/$(CODE_NAME)_$(OS_ID).tar.gz | awk '{print $$1}' > release/$(CODE_NAME)_$(OS_ID).sha256
	@echo "SHA256: $$(cat release/$(CODE_NAME)_$(OS_ID).sha256)"

install: release/$(CODE_NAME)_$(OS_ID)
	@echo "Installing the new $(CODE_NAME) binary..."
	@sudo cp release/$(CODE_NAME)_$(OS_ID) /usr/local/sbin/$(CODE_NAME)
	@sudo chmod 0755 /usr/local/sbin/$(CODE_NAME)
	@sudo chown 0:0 /usr/local/sbin/$(CODE_NAME)

buildonce:	notify-to-slack.go
	go build -ldflags "-w -s" -o $(CODE_NAME) $(CODE_NAME).go

clean:
	@rm -f notify-to-slack relase/$(CODE_NAME)_$(OS_ID) release/$(CODE_NAME)_$(OS_ID).tar.gz release/$(CODE_NAME)_$(OS_ID).sha256
