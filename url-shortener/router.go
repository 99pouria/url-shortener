package url_shortener

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strings"
)

func RunServer() {

	e := echo.New()

	e.GET("/open/:key", open)
	e.GET("/create/:url", createShortURL)

	e.Logger.Fatal(e.Start(":8080"))
}

func open(c echo.Context) error {
	longURL := getFromDB("http://localhost:8080/open/" + c.Param("key"))
	err := OpenUrl(longURL)
	if err != nil {
		fmt.Println(err)
	}
	err = c.String(http.StatusOK, longURL)
	return err
}

func createShortURL(c echo.Context) error {
	req := strings.Split(c.Request().RequestURI, "/create/")
	res, err := MapURLtoShorterURL(req[1])
	if err != nil {
		c.String(http.StatusOK, err.Error())
	} else {
		c.String(http.StatusOK, res)
	}
	return err
}
