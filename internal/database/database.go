package database

import (
	"github.com/go-redis/redis"
)

// DatabaseURL is an instance of URL database that maps long URLs with short ones.
type DatabaseURL struct {
	client *redis.Client
}

// NewClient creates new DB client with given database options.
func NewClient(address, password string, databaseNo int) DatabaseURL {
	return DatabaseURL{client: redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       databaseNo,
	})}
}

// Get returns value of matched key.
func (d DatabaseURL) Get(key string) (string, error) {
	return d.client.Get(key).Result()
}

// Add maps key to given value in database.
func (d DatabaseURL) Add(key, value string) error {
	return d.client.Set(key, value, 0).Err()
}

// KeyCount returns number of total keys stored in rdb.
func (d DatabaseURL) KeyCount() (int, error) {
	results, err := d.client.Keys("").Result()
	return len(results), err
}

// Close closes the client, releasing any open resources.
func (d DatabaseURL) Close() error {
	return d.client.Close()
}
