package main

import (
	"github.com/pooria1/url-shortener/database"
	"github.com/pooria1/url-shortener/server-handler"
)

func init() {
	database.InitializeDataBase()
}

func main() {
	server_handler.RunServer()
}
