package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/99pouria/url-shortener/internal/config"
	server_handler "github.com/99pouria/url-shortener/internal/server-handler"
	"golang.org/x/sys/unix"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, unix.SIGKILL)
	go func() {
		for sig := range c {
			fmt.Printf("signal %s received. closing...\n", sig.String())
			handle()
		}
	}()

	// loading config content
	if err := config.LoadConfig(); err != nil {
		log.Fatal(err)
	}

	s, err := server_handler.NewServer()
	if err != nil {
		log.Fatal(err)
	}
	handlers = append(handlers, s.Close)

	s.Run()
}

var handlers []func()

func handle() {
	for _, handle := range handlers {
		handle()
	}
}
