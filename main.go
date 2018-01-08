package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	negronilogrus "github.com/meatballhat/negroni-logrus"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"

	"github.com/henrywallace/mdpreview/server"
)

var (
	addr = flag.String("addr", ":8080", "address to serve preview")
	api  = flag.Bool("api", false, "whether to render via the Github API")
)

func main() {
	flag.Parse()
	log := logrus.New()

	if len(os.Args) < 2 {
		log.Fatal("path must be given")
	}
	path := os.Args[1]
	if filepath.Ext(path) != ".md" {
		log.Warnf("path %s doesn't look like a Markdown file", path)
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatalf("path %s does not exist", path)
	}

	s, err := server.New(path, log, !*api)
	if err != nil {
		log.Fatal(err)
	}
	h, err := s.Run()
	if err != nil {
		log.Fatal(err)
	}

	log.Info(fmt.Sprintf("Starting mdpreview server at http://localhost%s", *addr))
	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.Use(negronilogrus.NewMiddlewareFromLogger(log, "web"))
	n.UseHandler(h)
	if err := http.ListenAndServe(*addr, n); err != nil {
		log.Fatal(err)
	}
}
