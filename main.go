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

func (s *service) render() ([]byte, error) {
	f, err := os.Open(s.path)
	if err != nil {
		return nil, err
	}
	response, err := http.Post("https://api.github.com/markdown/raw", "text/plain", f)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (s *service) hRoot(w http.ResponseWriter, r *http.Request) {
	b, err := s.render()
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
