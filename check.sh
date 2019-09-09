#!/bin/sh

set -e

go install ./...
go mod tidy

go get honnef.co/go/tools/cmd/staticcheck
go get golang.org/x/lint/golint

set +e
go vet ./...
staticcheck ./... | rg -v server/asset.go
golint ./... | rg -v server/asset.go
set -e
