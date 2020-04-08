package url_shortener

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
)

var dataBase *redis.Client

func InitializeDataBase() {
	dataBase = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	UpdateUrlSet()
}

func getFromDB(key string) string {
	val, err := dataBase.Get(key).Result()
	if err == redis.Nil {
		addToDB(key, "")
		return ""
	}
	return val
}

func addToDB(key string, value string) {
	err := dataBase.Set(key, value, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func UpdateUrlSet() {
	s := getFromDB("urlShortener")
	if s == "" {
		return
	}
	err := json.Unmarshal([]byte(s), &UrlSet)
	if err != nil {
		fmt.Println(err)
	}
	for range UrlSet {
		seed++
	}
}

func updateDB() {
	dataBytes, _ := json.Marshal(UrlSet)
	addToDB("urlShortener", string(dataBytes))
}
