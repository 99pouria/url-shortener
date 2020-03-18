package url_shortener

import (
	"fmt"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
)

var seed int64
var urlSet map[string]string

func main() {

	urlSet = make(map[string]string)

	go http.ListenAndServe(":8080", nil)

	for {
		var choice int
		fmt.Print("[1] create a short link\n[2] show all links\n[other] quit\n\n")
		fmt.Scanf("%d", &choice)

		switch choice {
		case 1:
			var url string
			fmt.Print("type url: ")
			fmt.Scanf("%s", &url)
			fmt.Print("short lnk is: ", mapURLtoShorterURL(url), "\n")
		case 2:
			for key, value := range urlSet {
				fmt.Println("____________")
				fmt.Println("long link:", key)
				fmt.Print("short link: http://localhost:8080", value, "\n")
			}
			fmt.Println()
		default:
			break
		}

	}
}

func createShortAddress() string {
	seed++
	s := strconv.FormatInt(seed, 32)
	return fmt.Sprint("/", s)
}

func mapURLtoShorterURL(longUrl string) string {
	if !strings.Contains(longUrl, "http") {
		longUrl = fmt.Sprint("https://", longUrl)
	}

	urlSet[longUrl] = createShortAddress()

	http.HandleFunc(urlSet[longUrl], func(w http.ResponseWriter, r *http.Request) {
		openUrl(longUrl)
	})

	return fmt.Sprint("http://localhost:8080", urlSet[longUrl])
}

func openUrl(url string) {
	err := exec.Command("xdg-open", url).Run()
	if err != nil {
		fmt.Println(err)
	}
}
