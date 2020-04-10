package url_shortener

import (
	"github.com/go-redis/redis"
	"strconv"
)

var dataBase *redis.Client

func InitializeDataBase() error {
	dataBase = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if getFromDB("seed") == "" {
		return addToDB("seed", "0")
	} else {
		seed, _ = strconv.Atoi(getFromDB("seed"))
		return nil
	}
}

func getFromDB(key string) string {
	val, err := dataBase.Get(key).Result()
	if err == redis.Nil {
		addToDB(key, "")
		return ""
	}
	return val
}

func addToDB(key string, value string) error {
	err := dataBase.Set(key, value, 0).Err()
	return err
}
