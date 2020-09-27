package main

import (
	"flag"
	"github.com/hslam/websocket"
	"log"
	"net/http"
	_ "net/http/pprof"
)

var addr string

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
		Handler: websocket.Handler(Serve),
	}
	httpServer.ListenAndServe()
}

func Serve(conn *websocket.Conn) {
	for {
		msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		conn.WriteMessage(msg)
	}
	conn.Close()
}
