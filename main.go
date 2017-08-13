package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"mdpreview/server"
)

func main() {
	path := os.Args[1]
	s, err := server.New(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	h, err := s.Run()
	if err != nil {
		log.Fatal(err)
	}
	v := http.Server{
		Addr:    ":8080",
		Handler: h,
	}
	log.Fatal(v.ListenAndServe())
}
