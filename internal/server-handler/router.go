package server_handler

import (
	"fmt"
	"net/http"

	"github.com/99pouria/url-shortener/internal/config"
	"github.com/99pouria/url-shortener/internal/database"
	"github.com/99pouria/url-shortener/internal/shortener"
	"github.com/labstack/echo"
)

type Server struct {
	e  *echo.Echo
	db database.DatabaseURL
	sh *shortener.Shortener
}

func NewServer() (Server, error) {
	var (
		server Server
		cfg    = config.GetConfig()
	)

	// echo
	server.e = echo.New()

	// database
	server.db = database.NewClient(
		cfg.DatabaseConfig.Address,
		cfg.DatabaseConfig.Password,
		0,
	)

	keys, err := server.db.KeyCount()
	if err != nil {
		fmt.Println("e")
		return Server{}, err
	}

	// shortener
	server.sh = shortener.NewShortener(
		fmt.Sprintf(
			"%s:%d",
			cfg.ServerConfig.Hostname,
			cfg.ServerConfig.Port,
		),
		keys,
	)
	return server, nil
}

func (s Server) Run() {

	e := echo.New()

	e.GET("/:key", s.open)
	e.POST("/create", s.create)

	e.Logger.Fatal(e.Start(":8080"))
}

func (s Server) open(c echo.Context) error {
	longURL, err := s.db.Get(c.Param("key"))
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprintf("can't retrieve main URL: %s", err.Error()))
	}
	return c.Redirect(http.StatusMovedPermanently, longURL)
}

func (s Server) create(c echo.Context) error {
	longURL := c.FormValue("longURL")

	key, err := s.sh.GenerateNewKey()
	if err != nil {
		c.String(http.StatusOK, fmt.Sprintf("can't generate new key: %s", err.Error()))
	}

	s.db.Add(key, longURL)
	if err != nil {
		c.String(http.StatusOK, fmt.Sprintf("can't generate new key: %s", err.Error()))
	}
	return c.String(http.StatusOK, key)
}

func (s Server) Close() {
	s.db.Close()
	s.e.Close()
}
