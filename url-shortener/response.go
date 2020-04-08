package url_shortener

import (
	"fmt"
	"net/http"
	"strings"
)

func RunServer() {

	http.HandleFunc("/open/", func(writer http.ResponseWriter, request *http.Request) {
		req := strings.Split(request.RequestURI, "/open/")
		OpenUrl(UrlSet[req[1]])
		fmt.Fprintf(writer, UrlSet[req[1]])
	})

	http.HandleFunc("/create/", func(writer http.ResponseWriter, request *http.Request) {
		req := strings.Split(request.RequestURI, "/create/")
		res := MapURLtoShorterURL(req[1])
		fmt.Fprintf(writer, res)
	})

	http.HandleFunc("/showURLs", func(writer http.ResponseWriter, request *http.Request) {
		for key, value := range UrlSet {
			fmt.Fprintf(writer, "long link:\t%s\n", value)
			fmt.Fprintf(writer, "short link:\thttp://localhost:8080/open/%s\n\n", key)
		}
	})

	http.ListenAndServe(":8080", nil)
}
