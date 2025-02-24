.PHONY: setup-dev
setup-dev:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.57.2
	go install github.com/golang/mock/mockgen@v1.5.0

.PHONY: mock-gen
mock-gen:
	go generate ./...

.PHONY: lint
lint:
	@ golangci-lint run ./...

.PHONY: format
format:
	go fmt ./...

.PHONY: unit-test
unit-test:
	go test ./...