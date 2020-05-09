package generator

import (
	"github.com/go-redis/redis"
	"math/rand"
	"net/url"
	"os/exec"
	"url-shortener/database"
)

// letterRunes is list of characters that we use in short URL. It contains numbers, uppercase and lowercase letters
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func OpenUrl(url string) error {
	err := exec.Command("xdg-open", url).Run()
	return err
}

func CreateShortAddress() string {
	randString := make([]rune, 5)
	for i := range randString {
		randString[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(randString)
}

func MapURLtoShorterURL(longUrl string) (string, error) {

	_, err := url.ParseRequestURI(longUrl)
	if err != nil {
		return "", err
	}
	var shortUrl string
	for {
		shortUrl = "http://localhost:8080/open/" + CreateShortAddress()
		_, err := database.GetFromDB(shortUrl)
		if err == redis.Nil {
			break
		}
	}
	err = database.AddToDB(shortUrl, longUrl)
	return shortUrl, err
}
