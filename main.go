package main

import (
	"fmt"
	url_shortener "url-shortener/url-shortener"
)

func init() {
	err := url_shortener.InitializeDataBase()
	if err != nil {
		fmt.Println(err)
	}
	url_shortener.RunServer()
}

func main() {
}
