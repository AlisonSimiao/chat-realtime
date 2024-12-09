package server

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var port = ":8080"

func Init() {
	println("Starting websocket server on port " + *addr)

	http.HandleFunc("/ws", handleConnection)
	http.ListenAndServe(*addr, nil)
}

var addr = flag.String("addr", port, "http service address")

func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	println(conn.RemoteAddr().String() + " connected")
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
