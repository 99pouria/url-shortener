package build

import (
	"fmt"
	"url-shortener/build"
)

func main() {

	go build.RunServer()

	for {
		var choice int
		fmt.Print("[1] create a short link\n[2] show all links\n[other] quit\n")
		fmt.Scanf("%d", &choice)

		switch choice {
		case 1:
			var url string
			fmt.Print("type url: ")
			fmt.Scanf("%s", &url)
			fmt.Print("short link is: ", build.MapURLtoShorterURL(url), "\n\n")
		case 2:
			for key, value := range build.UrlSet {
				fmt.Println("____________")
				fmt.Println("long link:", key)
				fmt.Print("short link: http://localhost:8080", value, "\n")
			}
			fmt.Println()
		default:
			return
		}

	}
}
