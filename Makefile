.DEFAULT_GOAL := run

bins := ./website

fmt:
	go fmt .
.PHONY:fmt

lint: fmt
	golint .
.PHONY:lint

vet: lint
	go vet .
.PHONY:vet

run: vet
	go run .
.PHONY:run

build:
	go build .
.PHONY:build

clean:
	rm -fv $(bins)
.PHONY:clean
