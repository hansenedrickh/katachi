.PHONY: all
all: build fmt vet lint test

APP=katachi
ALL_PACKAGES=$(shell go list ./... | grep -v "vendor")
UNIT_TEST_PACKAGES=$(shell  go list ./... | grep -v "vendor")

DB_USER ?= postgres
DB_PASS ?= pass
DB_HOST ?= localhost
DB_PORT ?= 5432
DB_NAME ?= katachi

APP_EXECUTABLE="./out/$(APP)"

setup:
	GO111MODULE=off go get -u github.com/pressly/goose/cmd/goose
	GO111MODULE=off go get -u golang.org/x/lint/golint
	GO111MODULE=off go get -u github.com/axw/gocov/gocov
	GO111MODULE=off go get -u github.com/matm/gocov-html
	export GO111MODULE=on

compile:
	GO111MODULE=on go mod vendor
	templ generate
	mkdir -p out/
	go build -o $(APP_EXECUTABLE)

start:
	$(APP_EXECUTABLE) start

run: compile start

build: fmt vet lint compile

install:
	go install ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	@for p in $(UNIT_TEST_PACKAGES); do \
		echo "==> Linting $$p"; \
		golint $$p | { grep -vwE "exported (var|function|method|type|const) \S+ should have comment" || true; } \
	done

test:
	ENVIRONMENT=test go test -cover  ./... -coverprofile cover.out
	ENVIRONMENT=test go tool cover -func cover.out

test-cover-html:
	@echo "\nEXPORTING RESULTS TO COVERAGE.HTML..."
	gocov-html docs/cov.json > docs/coverage.html
	@echo 'TEST RESULTS EXPORTED TO DOCS/COVERAGE.HTML'

test-cov-report:
	@echo "\nGENERATING TEST REPORT."
	gocov report docs/cov.json

copy-config:
	cp application.yml.sample application.yml

protoc:
	protoc proto/*.proto --proto_path=./proto --go_out=contracts
	ls contracts/*.pb.go | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}'

migration-up:
	goose -dir migrations/ postgres "user=$(DB_USER) password=$(DB_PASS) dbname=$(DB_NAME) sslmode=disable" up

migration-down:
	goose -dir migrations/ postgres "user=$(DB_USER) password=$(DB_PASS) dbname=$(DB_NAME) sslmode=disable" down
