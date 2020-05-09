package main

import (
	"url-shortener/database"
	"url-shortener/server-handler"
)

func init() {
	database.InitializeDataBase()
}

func main() {
	server_handler.RunServer()
}
