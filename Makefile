mdpreview: static/github.css server/asset.go

.PHONY: server/asset.go
server/asset.go:
	go get -u github.com/jteeuwen/go-bindata/...
	go-bindata -pkg server -o server/asset.go static/...
	gofmt -w server/asset.go

.PHONY: static/github.css
static/github.css:
	npm install --global generate-github-markdown-css
	github-markdown-css > static/github.css