GOBIN ?= $$(go env GOPATH)/bin

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

# Pinned rather than @latest: newer releases of this tool require a newer Go
# than the one this module builds with.
GO_TEST_COVERAGE_VERSION ?= v2.12.0

.PHONY: install-go-test-coverage
install-go-test-coverage:
	go install github.com/vladopajic/go-test-coverage/v2@$(GO_TEST_COVERAGE_VERSION)

.PHONY: check-coverage
check-coverage: install-go-test-coverage
	go test ./... -coverprofile=./cover.out -covermode=atomic -coverpkg=./...
	${GOBIN}/go-test-coverage --config=./.testcoverage.yaml
