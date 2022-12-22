#!/usr/bin/env make
all : goadtouch goadtouch.exe
.PHONY : all

tidy:
	go mod tidy

get:
	go get -v

goadtouch:
	go build
	strip --strip-unneeded goadtouch
	upx goadtouch

goadtouch.exe:
	GOOS=windows GOARCH=amd64 go build
	strip --strip-unneeded goadtouch.exe
	upx goadtouch.exe

.PHONY : clean
clean:
	rm goadtouch goadtouch.exe
