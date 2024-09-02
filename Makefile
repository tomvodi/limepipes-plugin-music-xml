
.PHONY: test test-cover lint cover-html

build:
	go build -o ./limepipes-plugin-music-xml github.com/tomvodi/limepipes-plugin-music-xml/cmd/limepipes-plugin-music-xml

mocks:
	mockery

test:
	go test ./...

test-cover:
	go test ./... -coverprofile cover.out

lint:
	golangci-lint run

cover-html: test-cover
	go tool cover -html=cover.out
