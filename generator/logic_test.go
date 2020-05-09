package generator

import (
	"testing"
	"url-shortener/database"
)

func TestCreateShortAddress(t *testing.T) {
	expect := CreateShortAddress()
	if expect == "" {
		t.Error("function doesn't work and doesn't create path")
	}

}

func TestMapURLtoShorterURL(t *testing.T) {
	database.InitializeDataBase()
	url1 := "https://google.com/"
	url2 := "https://github.com/"
	shortURL1, err := MapURLtoShorterURL(url1)
	shortURL2, err := MapURLtoShorterURL(url2)
	if err != nil {
		t.Error(err)
	}
	if shortURL1 == "" || shortURL2 == "" {
		t.Error("result not expected")
	}
	database.DataBase.Del("http://localhost:8080/open/" + shortURL1)
	database.DataBase.Del("http://localhost:8080/open/" + shortURL2)
}
