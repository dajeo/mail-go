package main

import (
	"mail/config"
	"mail/db"
	"mail/server"
)

func main() {
	config.Init("dev")

	db.InitDB()

	server.Init()
}
