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

db-reset:
  rm -f db/data.db
  mkdir -p db
