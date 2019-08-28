
PKGS := $(shell echo $(PKGS_AND_MOCKS) | tr ' ' '\n' | grep -v /mock$)

local:
	export GO111MODULE=on && go build && ./go_api

coverage:
	@go test ./... -coverprofile=coverage.out -parallel 4
	@go tool cover -html=coverage.out

test:
	@go test ./... -parallel 4