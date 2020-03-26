package build

import (
	"fmt"
	"github.com/labstack/echo"
	"os/exec"
	"strconv"
	"strings"
)

var seed int64
var UrlSet = make(map[string]string)

func RunServer(e *echo.Echo) {
	e.Logger.Fatal(e.Start(":8080"))
}

func CreateShortAddress() string {
	seed++
	s := strconv.FormatInt(seed, 32)
	return fmt.Sprint(s)
}

func MapURLtoShorterURL(longUrl string, e *echo.Echo) string {
	if !strings.Contains(longUrl, "http") {
		longUrl = fmt.Sprint("https://", longUrl)
	}

	UrlSet[longUrl] = CreateShortAddress()

	e.GET("/:shortUrl", func(context echo.Context) error {
		return OpenUrl(longUrl)
	})

	return fmt.Sprint("http://localhost:8080/", UrlSet[longUrl])
}

func OpenUrl(url string) error {
	err := exec.Command("xdg-open", url).Run()
	return err
}
