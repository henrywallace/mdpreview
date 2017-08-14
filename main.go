package main

import (
	"fmt"
	"net/http"
	"os"

	negronilogrus "github.com/meatballhat/negroni-logrus"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"

	"github.com/henrywallace/mdpreview/server"
)

func main() {
	log := logrus.New()

	path := os.Args[1]
	s, err := server.New(path, log)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	h, err := s.Run()
	if err != nil {
		log.Fatal(err)
	}

	addr := ":8080"
	log.Info(fmt.Sprintf("Starting mdpreview server at http://localhost%s", addr))
	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.Use(negronilogrus.NewMiddlewareFromLogger(log, "web"))
	n.UseHandler(h)
	if err := http.ListenAndServe(addr, n); err != nil {
		log.Fatal(err)
	}
}
