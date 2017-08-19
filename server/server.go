package server

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/russross/blackfriday"
	"github.com/sirupsen/logrus"
)

// Server serves a HTML rendered Markdown preview of a Markdown file specified
// at path. Whenever the path is written to, the rendering will update
// dynamically.
type Server struct {
	path          string
	indexTemplate *template.Template
	upgrader      websocket.Upgrader
	log           *logrus.Logger
	renderLocally bool
}

// New creates a new Server given some markdown path.
func New(path string, log *logrus.Logger, renderLocally bool) (*Server, error) {
	indexData, err := Asset("static/index.html")
	if err != nil {
		log.Fatal("index not found")
	}
	indexTemplate := template.Must(template.New("index").Parse(string(indexData)))
	if err != nil {
		log.Fatal("index template parse failed")
	}

	return &Server{
		path:          path,
		log:           log,
		indexTemplate: indexTemplate,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		renderLocally: renderLocally,
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
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(indexBuf.Bytes())
}

func (s *Server) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	ws, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			s.log.WithError(err)
		}
		return
	}

	go s.writer(ws)
	s.reader(ws)
}

func (s *Server) render() ([]byte, error) {
	switch s.renderLocally {
	case true:
		input, err := ioutil.ReadFile(s.path)
		if err != nil {
			return nil, err
		}
		b := blackfriday.MarkdownCommon(input)
		return b, nil
	case false:
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
	default:
		panic("unreachable")
	}
}

func (s *Server) watcher(changes chan<- struct{}) {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		s.log.Fatal(err)
	}
	defer w.Close()

	changes <- struct{}{} // for initial display
	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-w.Events:
				s.log.WithFields(logrus.Fields{
					"file":  event.Name,
					"event": event.Op,
				}).Info("changed file")
				if event.Op&fsnotify.Write == fsnotify.Write {
					changes <- struct{}{}
				}
			case err := <-w.Errors:
				s.log.WithError(err)
			}
		}
	}()

	err = w.Add(s.path)
	if err != nil {
		s.log.Fatal(err)
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
			rendered, err := s.render()
			if err != nil {
				s.log.Error(err)
			}
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
