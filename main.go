package main

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/HsiaoCz/search-engine/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		slog.Error("load the env error", "err", err)
		return
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = ":4001"
	}

	app := fiber.New(fiber.Config{
		IdleTimeout: 5 * time.Second,
	})

	app.Use(compress.New())

	// set routers
	routers.SetRouters(app)
	// start server and listen for shutdown

	go func() {
		if err := app.Listen(port); err != nil {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c // Block  the main thread until interupted

	app.Shutdown()

	slog.Info("the server is shuting down")
}
