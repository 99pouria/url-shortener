package build_test

import (
	"fmt"
	"strconv"
	"testing"
	"url-shortener/build"
)

func TestCreateShortAddress(t *testing.T) {
	go build.RunServer()

	for i := 1; i < 100; i++ {
		expect := build.CreateShortAddress()
		got := fmt.Sprint("/", strconv.FormatInt(int64(i), 32))
		if expect != got {
			t.Error("expected ", expect, ", got ", got)
		}
	}
}

func TestOpenUrl(t *testing.T) {
	err := build.OpenUrl("http://google.com/")
	if err != nil {
		t.Error(err)
	}
}

func TestMapURLtoShorterURL(t *testing.T) {
	go build.RunServer()
	url1 := "https://google.com/"
	url2 := "https://github.com/"
	build.MapURLtoShorterURL(url1)
	build.MapURLtoShorterURL(url2)

	if build.UrlSet[url1] != "/34" || build.UrlSet[url2] != "/35" {
		t.Error("result not expected, urlSet:\n", build.UrlSet)
	}
}
