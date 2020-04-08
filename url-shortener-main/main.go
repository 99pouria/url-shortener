package main

import (
	"url-shortener/url-shortener"
)

func main() {

	url_shortener.InitializeDataBase()
	url_shortener.RunServer()

}
