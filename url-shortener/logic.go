package url_shortener

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

var seed int64
var UrlSet = make(map[string]string)

func OpenUrl(url string) error {
	err := exec.Command("xdg-open", url).Run()
	return err
}

func CreateShortAddress() string {
	seed++
	s := strconv.FormatInt(seed, 32)
	return fmt.Sprint(s)
}

func MapURLtoShorterURL(longUrl string) string {
	if !strings.Contains(longUrl, "http") {
		longUrl = fmt.Sprint("https://", longUrl)
	}

	shortString := CreateShortAddress()
	UrlSet[shortString] = longUrl
	updateDB()

	return fmt.Sprint("http://localhost:8080/open/", shortString)
}
