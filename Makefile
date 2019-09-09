mdpreview: static/github.css server/asset.go

.PHONY: server/asset.go
server/asset.go:
	GO111MODULES=off go get -u github.com/jteeuwen/go-bindata/...
	go-bindata -pkg server -o server/asset.go static/...
	gofmt -w server/asset.go

.PHONY: static/github.css
static/github.css:
	npm install --global generate-github-markdown-css
	github-markdown-css > static/github.css
	go get github.com/tdewolff/minify/cmd/minify
	minify -o static/github.css static/github.css
