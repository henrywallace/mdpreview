mdpreview: server/asset.go

.PHONY: server/asset.go
server/asset.go:
	go get -u github.com/jteeuwen/go-bindata/...
	go-bindata -pkg server -o server/asset.go static/...
	gofmt -w server/asset.go