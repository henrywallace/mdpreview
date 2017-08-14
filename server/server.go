package server

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

// Server serves a HTML rendered Markdown preview of a Markdown file specified
// at path. Whenever the path is written to, the rendering will update
// dynamically.
type Server struct {
	path          string
	indexTemplate *template.Template
	upgrader      websocket.Upgrader
}

// New creates a new Server given some markdown path.
func New(path string) (*Server, error) {
	indexData, err := Asset("static/index.html")
	if err != nil {
		panic("index not found")
	}
	indexTemplate := template.Must(template.New("index").Parse(string(indexData)))
	if err != nil {
		panic("index template parse failed")
	}

	return &Server{
		path:          path,
		indexTemplate: indexTemplate,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}, nil
}

// Run returns handlers to run the server.
func (s *Server) Run() (http.Handler, error) {
	return s.setupHandlers(), nil
}

func (s *Server) setupHandlers() http.Handler {
	staticFileHandler := http.FileServer(&assetfs.AssetFS{
		Asset:    Asset,
		AssetDir: AssetDir,
		Prefix:   "static",
	})

	r := mux.NewRouter()
	r.HandleFunc("/", s.handleIndex).Methods("GET")
	r.HandleFunc("/ws", s.handleWebSocket).Methods("GET")

	r.PathPrefix("/preview.js").Handler(staticFileHandler).Methods("GET")
	r.PathPrefix("/favicon.ico").Handler(staticFileHandler).Methods("GET")

	return r
}

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	indexBuf := new(bytes.Buffer)
	err := s.indexTemplate.Execute(indexBuf, map[string]interface{}{
		"path": filepath.Base(s.path),
	})
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	w.Write(indexBuf.Bytes())
}

func (s *Server) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	ws, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Println(err)
		}
		return
	}

	go s.writer(ws)
	s.reader(ws)
}

func (s *Server) render() ([]byte, error) {
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

func (s *Server) watcher(changes chan<- struct{}) {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer w.Close()

	changes <- struct{}{} // for initial display
	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-w.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					changes <- struct{}{}
				}
			case err := <-w.Errors:
				log.Println("error:", err)
			}
		}
	}()

	err = w.Add(s.path)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func (s *Server) writer(ws *websocket.Conn) {
	defer func() {
		ws.Close()
	}()

	changes := make(chan struct{})
	go s.watcher(changes)
	for {
		select {
		case <-changes:
			rendered, _ := s.render()
			ws.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := ws.WriteMessage(websocket.TextMessage, rendered); err != nil {
				return
			}
		case <-time.After(2 * time.Second):
			ws.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func (s *Server) reader(ws *websocket.Conn) {
	defer ws.Close()
	ws.SetReadLimit(512)
	ws.SetReadDeadline(time.Now().Add(60 * time.Second))
	ws.SetPongHandler(func(string) error {
		ws.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})
	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			break
		}
	}
}
