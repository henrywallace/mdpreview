package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type service struct {
	path string
}

func (s *service) hRoot(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open(s.path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	response, err := http.Post("https://api.github.com/markdown/raw", "text/plain", f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "text/html")
	w.Write(b)
}

func main() {
	s := service{
		path: os.Args[1],
	}
	http.HandleFunc("/", s.hRoot)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
