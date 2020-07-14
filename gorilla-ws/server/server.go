package main

import (
	"flag"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	_ "net/http/pprof"
)

var addr string
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024 * 64,
	WriteBufferSize: 1024 * 64,
}

func init() {
	flag.StringVar(&addr, "addr", ":9999", "-addr=:9999")
	flag.Parse()
}
func main() {
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()
	httpServer := &http.Server{
		Addr:    addr,
		Handler: new(Handler),
	}
	httpServer.ListenAndServe()
}

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.Header.Del("Origin")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		conn.WriteMessage(websocket.BinaryMessage, msg)
	}
}
