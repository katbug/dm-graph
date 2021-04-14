.PHONY: build
build:
	@go build

.PHONY: serve
serve: 
	@go run server.go

