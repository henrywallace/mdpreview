package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/websocket"
)

const indexHTML = `
<!DOCTYPE html>
<html lang="en">
    <head>
        <title>Markdown Preview</title>
	</head>
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<link rel="stylesheet" href="https://rawgit.com/sindresorhus/github-markdown-css/gh-pages/github-markdown.css">
	<style>
		.markdown-body {
			box-sizing: border-box;
			min-width: 200px;
			max-width: 980px;
			margin: 0 auto;
			padding: 45px;
		}
		@media (max-width: 767px) {
			.markdown-body {
				padding: 15px;
			}
		}
	</style>
    <body>
        <article id="preview" class="markdown-body" type=html>{{.Data}}</article>
        <script type="text/javascript">
            (function() {
                var preview = document.getElementById("preview");
                var conn = new WebSocket("ws://{{.Host}}/ws");
                conn.onclose = function(evt) {
                    preview.innerHTML = 'Connection closed';
                }
                conn.onmessage = function(event) {
                    preview.innerHTML = event.data;
                }
            })();
        </script>
    </body>
</html>
`

type service struct {
	path          string
	indexTemplate *template.Template
	upgrader      websocket.Upgrader
}

func (s *service) serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	rendered, _ := s.render()
	var v = struct {
		Host string
		Data template.HTML
	}{
		r.Host,
		template.HTML(rendered),
	}
	s.indexTemplate.Execute(w, &v)
}

func (s *service) serveWs(w http.ResponseWriter, r *http.Request) {
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

func (s *service) watcher(changes chan<- struct{}) {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer w.Close()

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

func (s *service) writer(ws *websocket.Conn) {
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

func (s *service) reader(ws *websocket.Conn) {
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

func main() {
	s := service{
		path:          os.Args[1],
		indexTemplate: template.Must(template.New("").Parse(indexHTML)),
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}
	http.HandleFunc("/", s.serveHome)
	http.HandleFunc("/ws", s.serveWs)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
