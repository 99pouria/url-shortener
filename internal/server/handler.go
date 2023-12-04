package server

import (
	"fmt"
	"net/http"
	"net/url"
	"sync/atomic"

	"github.com/99pouria/url-shortener/internal/config"
	"github.com/99pouria/url-shortener/internal/database"
	"github.com/99pouria/url-shortener/internal/shortener"
	"github.com/labstack/echo"
)

// Server is instance of server
type Server struct {
	echoInstance *echo.Echo
	db           database.DatabaseURL
	shortenerAPI *shortener.Shortener
	address      string      // address of server
	isStart      atomic.Bool // isStart shows if server is running or not
}

// NewServer creates new connection to redis database and starts serving address (which has been set in
// config file).
//
// This server accepts a POST method with 'longURL' form value and returns a short URL that if you call
// the URL it redirects to original URL (long URL).
func NewServer() (*Server, error) {
	var (
		server Server
		cfg    = config.GetConfig()
	)

	server.address = fmt.Sprintf("%s:%d", cfg.ServerConfig.Hostname, cfg.ServerConfig.Port)
	server.isStart.Store(false)

	// echo
	server.echoInstance = echo.New()

	// database
	server.db = database.NewClient(
		cfg.DatabaseConfig.Address,
		cfg.DatabaseConfig.Password,
		0,
	)

	keys, err := server.db.KeyCount()
	if err != nil {
		return nil, err
	}

	// shortener
	server.shortenerAPI = shortener.NewShortener(keys)
	return &server, nil
}

// Run starts serving clients
//
// Do not forget to defer 'Close' function
func (s *Server) Run() error {
	if !s.isStart.CompareAndSwap(false, true) {
		return fmt.Errorf("server is already running")
	}
	defer s.isStart.Store(false)

	s.echoInstance.GET("/:key", s.open)
	s.echoInstance.POST("/create", s.create)
	return s.echoInstance.Start(s.address)
}

// open redirects given url to original URL
func (s *Server) open(c echo.Context) error {
	longURL, err := s.db.Get(c.Param("key"))
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("original URL not found: %s", err.Error()))
	}

	return c.Redirect(http.StatusMovedPermanently, longURL)
}

// create maps original URL to a short and unique URL and sends short URL as response
func (s *Server) create(c echo.Context) error {
	key, err := s.shortenerAPI.GenerateNewKey()
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("can't generate new URL: %s", err.Error()))
	}

	s.db.Add(key, c.FormValue("longURL"))
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("can't add short URL to database: %s", err.Error()))
	}

	shortURL, err := url.JoinPath("https://", s.address, key)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("can't create short url: %s", err.Error()))
	}
	return c.String(http.StatusOK, shortURL)
}

// Close closes the connection to database and http serving connection
func (s *Server) Close() {
	s.db.Close()
	s.echoInstance.Close()
}
