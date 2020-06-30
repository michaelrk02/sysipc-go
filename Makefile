GO = go
BUILD = $(GO) build
REPO = github.com/michaelrk02/sysipc-go
OUT = -o build

all: test-server test-client
.PHONY: all

test-server:
	$(BUILD) $(OUT)/test-server $(REPO)/test-server
.PHONY: test-server

test-client:
	$(BUILD) $(OUT)/test-client $(REPO)/test-client
.PHONY: test-client

