package url_shortener

import (
	"fmt"
	"github.com/labstack/echo"
	"strconv"
	"testing"
)

var e *echo.Echo

func TestCreateShortAddress(t *testing.T) {
	go RunServer()

	for i := 1; i < 100; i++ {
		expect := CreateShortAddress()
		got := fmt.Sprint("/", strconv.FormatInt(int64(i), 32))
		if expect != got {
			t.Error("expected ", expect, ", got ", got)
		}
	}
}

func TestMapURLtoShorterURL(t *testing.T) {
	go RunServer()
	url1 := "https://google.com/"
	url2 := "https://github.com/"
	MapURLtoShorterURL(url1)
	MapURLtoShorterURL(url2)

	if UrlSet[url1] != "/34" || UrlSet[url2] != "/35" {
		t.Error("result not expected, urlSet:\n", UrlSet)
	}
}
