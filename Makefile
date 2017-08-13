mdpreview: server/asset.go

.PHONY: server/asset.go
server/asset.go:
	go-bindata -pkg server -o server/asset.go static/...
	gofmt -w server/asset.go