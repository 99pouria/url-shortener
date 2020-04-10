package url_shortener

import (
	"fmt"
	"net/http"
	"strings"
)

func RunServer() {

	http.HandleFunc("/open/", func(writer http.ResponseWriter, request *http.Request) {
		longURL := getFromDB("http://localhost:8080" + request.RequestURI)
		err := OpenUrl(longURL)
		if err != nil {
			fmt.Println(err)
		}
		_, err = fmt.Fprintf(writer, longURL)
		if err != nil {
			fmt.Println(err)
		}
	})

	http.HandleFunc("/create/", func(writer http.ResponseWriter, request *http.Request) {
		req := strings.Split(request.RequestURI, "/create/")
		res := MapURLtoShorterURL(req[1])
		_, err := fmt.Fprintf(writer, res)
		if err != nil {
			fmt.Println(err)
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println()
	}
}

func init() {
	err := InitializeDataBase()
	if err != nil {
		fmt.Println(err)
	}
	RunServer()
}

func Run() {
}
