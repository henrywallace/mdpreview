package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	negronilogrus "github.com/meatballhat/negroni-logrus"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"

	"github.com/henrywallace/mdpreview/server"
)

var (
	addr  = flag.String("addr", ":8080", "address to serve preview like :8080 or 0.0.0.0:7000")
	api   = flag.Bool("api", false, "whether to render via the Github API")
	debug = flag.Bool("debug", false, "debug logging")
)

func main() {
	flag.Parse()

	log := logrus.New()
	if *debug {
		log.SetLevel(logrus.DebugLevel)
	}

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

	if strings.HasPrefix(*addr, ":") {
		*addr = fmt.Sprintf("127.0.0.1%s", *addr)
	}

	log.Info(fmt.Sprintf("Starting mdpreview server at %s", *addr))
	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.Use(negronilogrus.NewMiddlewareFromLogger(log, "web"))
	n.UseHandler(h)
	if err := http.ListenAndServe(*addr, n); err != nil {
		log.Fatal(err)
	}
}
