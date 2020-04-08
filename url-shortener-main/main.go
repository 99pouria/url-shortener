package main

import (
	"url-shortener/build"
)

func main() {

	build.InitializeDataBase()
	build.NewDateBaseClient()
	build.UpdateUrlSet()
	build.RunServer()

}
