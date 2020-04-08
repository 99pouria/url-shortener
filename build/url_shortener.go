package build

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var seed int64
var UrlSet = make(map[string]string)
var dataBase *redis.Client

func RunServer() {

	http.HandleFunc("/open/", func(writer http.ResponseWriter, request *http.Request) {
		req := strings.Split(request.RequestURI, "/open/")
		OpenUrl(UrlSet[req[1]])
		fmt.Fprintf(writer, UrlSet[req[1]])
	})

	http.HandleFunc("/create/", func(writer http.ResponseWriter, request *http.Request) {
		req := strings.Split(request.RequestURI, "/create/")
		res := MapURLtoShorterURL(req[1])
		fmt.Fprintf(writer, res)
	})

	http.HandleFunc("/showURLs", func(writer http.ResponseWriter, request *http.Request) {
		for key, value := range UrlSet {
			fmt.Fprintf(writer, "long link:\t%s\n", value)
			fmt.Fprintf(writer, "short link:\thttp://localhost:8080/open/%s\n\n", key)
		}
	})

	http.ListenAndServe(":8080", nil)
}

func NewDateBaseClient() {
	dataBase := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := dataBase.Ping().Result()
	if err != nil {
		fmt.Println(err)
	}
}

func InitializeDataBase() {
	dataBase = redis.NewClient(&redis.Options{
		Addr:         ":6379",
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})
}

func OpenUrl(url string) error {
	err := exec.Command("xdg-open", url).Run()
	return err
}

func CreateShortAddress() string {
	seed++
	s := strconv.FormatInt(seed, 32)
	return fmt.Sprint(s)
}

func MapURLtoShorterURL(longUrl string) string {
	if !strings.Contains(longUrl, "http") {
		longUrl = fmt.Sprint("https://", longUrl)
	}

	shortString := CreateShortAddress()
	UrlSet[shortString] = longUrl
	updateDB()

	return fmt.Sprint("http://localhost:8080/open/", shortString)
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
}

func updateDB() {
	dataBytes, _ := json.Marshal(UrlSet)
	addToDB("urlShortener", string(dataBytes))
}
