name := $(shell basename ${PWD})

all: help

.PHONY: help
help: Makefile
	@echo
	@echo " Welcome to ${name}! 🎉"
	@echo
	@echo " Choose a make command to run:"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

## init: Initialize project
.PHONY: init
init:
	cd src && go mod tidy
	go install github.com/cosmtrek/air@latest
	alias air='${HOME}/go/bin/air'

## test: Test packages
.PHONY: test
test:
	go clean -testcache
	go test -v ./src/containers/...

## build: Build app binary
.PHONY: build
build: test
	cd src/cmd && go build -o ./.air/main ./

## docker-build: Build app docker image
.PHONY: docker-build
docker-build: test
	GOPROXY=direct docker buildx build -t ${name} .

## docker-run: Run app in docker container
.PHONY: docker-run
docker-run:
	docker run -it --rm -p 8080:8080 ${name}

## start: Build and run local project
.PHONY: start
start: build
	air

## css: Build tailwindcss
.PHONY: css
css:
	cd src/configs && tailwindcss -i ../web/css/input.css -o ../web/css/output.css --minify

## css-watch: Build tailwindcss in watch mode
.PHONY: css-watch
css-watch:
	cd src/configs && tailwindcss -i ../web/css/input.css -o ../web/css/output.css --watch
