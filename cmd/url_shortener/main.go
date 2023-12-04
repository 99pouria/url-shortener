package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/99pouria/url-shortener/internal/config"
	"github.com/99pouria/url-shortener/internal/server"
	"golang.org/x/sys/unix"
)

func main() {
	// handling ctrl+c and SIGKILL signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, unix.SIGKILL)
	go func() {
		for sig := range c {
			fmt.Printf("signal %s received. closing...\n", sig.String())
			handle()
			os.Exit(0)
		}
	}()
	defer handle()

	// loading config content
	if err := config.LoadConfig(); err != nil {
		log.Fatal(err)
	}

	// creating new server instance
	s, err := server.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	// lets run!
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
	handlers = append(handlers, s.Close) // add close function of server to handlers
}

// handlers containts all handler which must call before finishing process
var handlers []func()

// handle calls all handler functions
func handle() {
	for _, handle := range handlers {
		handle()
	}
}
