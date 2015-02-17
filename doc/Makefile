#!/bin/bash
.PHONY: detect_os test Linux Darwin clean

export GOPATH=$(PWD)
UNAME_S = $(shell uname -s)

ifeq ($(UNAME_S), Linux)
	export CGO_LDFLAGS=-L $(PWD)/Release -lcef
else #ifeq ($(UNAME_S), Darwin)
	export CGO_ENABLED=1
	export CGO_LDFLAGS=-F$(PWD)/Release/tmp -framework Cocoa -framework cef
endif

detect_os:
	make $(UNAME_S)

test:
	clear
	cd ./src/github.com/24hours/chrome ; pwd ; export CGO_LDFLAGS="-L `pwd`/Release -lcef" ; go test -o Release/chrome.test -ldflags "-r ./Release" ; cd ../../../../

Linux:
	clear
	go install github.com/24hours/chrome
	go build -ldflags "-r ." -o Release/cef2go main_linux.go
	
Darwin:
	clear
	go install github.com/24hours/chrome
	@# Required for linking. Go doesn't allow framework name
	@# to contain spaces, so we're making a copy of the framework
	@# without spaces.
	@if [ ! -d Release/tmp ]; then \
		echo Copying CEF framework directory to Release/tmp ;\
		mkdir -p Release/tmp ;\
		cp -rf Release/cef2go.app/Contents/Frameworks/Chromium\ Embedded\ Framework.framework Release/tmp/cef.framework ;\
		mv Release/tmp/cef.framework/Chromium\ Embedded\ Framework Release/tmp/cef.framework/cef ;\
	fi
	@# CEF requires specific app bundle / directory structure
	@# on OSX, but Go doesn't allow for such thing when 
	@# running test. So turning off test.
	@# go test -ldflags "-r $(PWD)/Release" src/tests/cef_test.go
	rm -f Release/cef2go.app/Contents/MacOS/cef2go
	go build -ldflags "-r ." -o Release/cef2go.app/Contents/MacOS/cef2go main_darwin.go
	install_name_tool -change @executable_path/Chromium\ Embedded\ Framework @executable_path/../Frameworks/Chromium\ Embedded\ Framework.framework/Chromium\ Embedded\ Framework Release/cef2go.app/Contents/MacOS/cef2go
	
clean:
	clear
	go clean -i cef
