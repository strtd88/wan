#
#
# justfile
#

help:
    @just --list

#
#
#

default:
  just help

run:
  go run main.go

build:
  go build -o wan

install:
  go install .

test:
  go test ./test -v

deps:
  go get github.com/spf13/cobra@latest
  go get github.com/mattn/go-sqlite3@latest
  go get github.com/stretchr/testify@latest

db-shell:
  sqlite3 db/data.db

db-deploy:
  sqitch deploy

db-revert:
  sqitch revert HEAD^

db-reset:
  rm -f db/data.db
  mkdir -p db
  just db-deploy

lint:
  golangci-lint run --config .golangci.yml

test-coverage:
  go test ./test -v -coverprofile=coverage.out
  go tool cover -html=coverage.out -o coverage.html
