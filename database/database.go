package database

import (
	"github.com/go-redis/redis"
)

var DataBase *redis.Client

func init() {
	DataBase = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

}

func GetFromDB(key string) (string, error) {
	val, err := DataBase.Get(key).Result()
	return val, err
}

func AddToDB(key string, value string) error {
	err := DataBase.Set(key, value, 0).Err()
	return err
}
