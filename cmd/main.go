package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sotnasr/url-shortener/internal/cache"
	"github.com/sotnasr/url-shortener/internal/delivery/http"
)

func main() {
	// Setup cache
	cache, err := cache.NewRedisCache(os.Getenv("REDIS_HOSTNAME"), os.Getenv("REDIS_PASSWORD"), os.Getenv("REDIS_PORT"))
	if err != nil {
		log.Fatal("error was found while connect on redis: ", err)
	}

	// Setup route engine & middleware
	e := echo.New()

	http.NewShortenerHandler(e, cache)

	go func() {
		if err := e.Start(":9000"); err != nil {
			log.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
