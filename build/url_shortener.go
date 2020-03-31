package build

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var seed int64
var UrlSet = make(map[string]string)
var dataBase *redis.Client

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
	updateDB()

	e.GET("/:shortUrl", func(context echo.Context) error {
		return OpenUrl(longUrl)
	})

	return fmt.Sprint("http://localhost:8080/", UrlSet[longUrl])
}

func OpenUrl(url string) error {
	err := exec.Command("xdg-open", url).Run()
	return err
}

func Initialize() {
	dataBase = redis.NewClient(&redis.Options{
		Addr:         ":6379",
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})
}

func NewClient() {
	dataBase := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // use default Addr
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	_, err := dataBase.Ping().Result()
	if err != nil {
		fmt.Println(err)
	}
}

func getKey(key string) string {
	val, err := dataBase.Get(key).Result()
	if err == redis.Nil {
		fmt.Println("DB is empty")
		setKey(key, "")
		return ""
	}
	return val
}

func setKey(key string, value string) {
	err := dataBase.Set(key, value, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func GetUrlSetFromDB() {
	s := getKey("urlShortener")
	if s == "" {
		return
	}
	err := json.Unmarshal([]byte(s), &UrlSet)
	if err != nil {
		fmt.Println(err)
	}
}

func updateDB() {
	dataBytes, _ := json.Marshal(UrlSet)
	setKey("urlShortener", string(dataBytes))
}
