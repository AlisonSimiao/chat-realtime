package main

import (
	"github.com/AlisonSimiao/chat-realtime/database"
	"github.com/AlisonSimiao/chat-realtime/server"
)

func main() {
	database.InitDB()
	server.Init()
}
