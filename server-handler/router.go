package server_handler

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/pooria1/url-shortener/database"
	"github.com/pooria1/url-shortener/generator"
	"net/http"
)

func RunServer() {

	e := echo.New()

	e.GET("/open/:key", open)
	e.POST("/create", create)

	e.Logger.Fatal(e.Start(":8080"))
}

func open(c echo.Context) error {
	longURL, err := database.GetFromDB("http://localhost:8080/open/" + c.Param("key"))
	if err != nil {
		fmt.Println(err)
	}
	err = generator.OpenUrl(longURL)
	if err != nil {
		fmt.Println(err)
	}
	err = c.String(http.StatusOK, longURL)
	return err
}

func create(c echo.Context) error {
	longURL := c.FormValue("longURL")
	res, err := generator.MapURLtoShorterURL(longURL)
	if err != nil {
		c.String(http.StatusOK, err.Error())
	} else {
		c.String(http.StatusOK, res)
	}
	return err
}
