package url_shortener

import (
	"fmt"
	"net/url"
	"os/exec"
	"strconv"
	"strings"
)

var seed int
var UrlSet = make(map[string]string)

func OpenUrl(url string) error {
	err := exec.Command("xdg-open", url).Run()
	return err
}

func CreateShortAddress() string {
	seed++
	err := addToDB("seed", fmt.Sprint(seed))
	if err != nil {
		fmt.Println(err)
	}
	return strconv.FormatInt(int64(seed), 32)
}

func MapURLtoShorterURL(longUrl string) (string, error) {

	if !strings.Contains(longUrl, "http") {
		longUrl = "https://" + longUrl
	}
	_, err := url.ParseRequestURI(longUrl)
	if err != nil {
		return "", err
	}

	shortUrl := "http://localhost:8080/open/" + CreateShortAddress()
	err = addToDB(shortUrl, longUrl)

	if err != nil {
		fmt.Println(err)
	}

	return shortUrl, err
}
