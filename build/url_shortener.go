package build

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
)

var seed int64
var UrlSet = make(map[string]string)

func RunServer() {
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateShortAddress() string {
	seed++
	s := strconv.FormatInt(seed, 32)
	return fmt.Sprint("/", s)
}

func MapURLtoShorterURL(longUrl string) string {
	if !strings.Contains(longUrl, "http") {
		longUrl = fmt.Sprint("https://", longUrl)
	}

	UrlSet[longUrl] = CreateShortAddress()

	http.HandleFunc(UrlSet[longUrl], func(w http.ResponseWriter, r *http.Request) {
		OpenUrl(longUrl)
	})

	return fmt.Sprint("http://localhost:8080", UrlSet[longUrl])
}

func OpenUrl(url string) error {
	err := exec.Command("xdg-open", url).Run()
	return err
}
