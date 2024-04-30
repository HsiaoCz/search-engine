package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/HsiaoCz/search-engine/conf"
	"github.com/HsiaoCz/search-engine/routers"
	"github.com/HsiaoCz/search-engine/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if err := godotenv.Load(); err != nil {
		slog.Error("load the env error", "err", err)
		return
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(conf.GetMongoUrl("DATABASE_URL")))
	if err != nil {
		slog.Error("connect mongo db error", "err", err)
		return
	}

	var (
		port      = conf.GetPort("PORT")
		dbname    = conf.GetMongoDBName("DBNAME")
		userColl  = conf.GetUserColl("USERCOLL")
		userStore = storage.NewMongoUserStore(client, dbname, userColl)
		store     = &storage.Store{User: userStore}
	)

	app := fiber.New(fiber.Config{
		IdleTimeout: 5 * time.Second,
	})

	app.Use(compress.New())

	// set routers
	routers.SetRouters(app, store)
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
