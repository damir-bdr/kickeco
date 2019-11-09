GOCMD=go
GOBUILD=$(GOCMD) build

all: build
build:
	GO111MODULE=on $(GOBUILD) -o kick-invoker main.go
clean:
	rm -f kick-invoker
