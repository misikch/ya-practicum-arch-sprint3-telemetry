# Makefile

# Укажите путь к бинарному файлу golangci-lint
GOLANGCI_LINT := $(shell which golangci-lint)

.PHONY: all
all: generate lint test build

.PHONY: generate
generate:
	go generate ./...

.PHONY: lint
lint: $(GOLANGCI_LINT)
	$(GOLANGCI_LINT) run

.PHONY: test
test:
	go test ./...

.PHONY: build
build: build-service build-worker

.PHONY: build-service
build-service:
	go build -o ./build/service ./cmd/service/main.go

.PHONY: build-worker
build-worker:
	go build -o ./build/worker ./cmd/worker/main.go

# Установка golangci-lint, если он не установлен
$(GOLANGCI_LINT):
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.54.2
