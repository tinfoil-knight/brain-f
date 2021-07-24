run:
	@echo "> Starting server"
	go run .

format:
	@echo "> Formatting the source"
	gofmt -d -e .

build:
	@echo "> Building binary"
	go build -o bf .

.PHONY: run format build