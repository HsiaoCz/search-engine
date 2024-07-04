package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/HsiaoCz/search-engine/routers"
	"github.com/HsiaoCz/search-engine/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("DATABASE_URL")))
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := client.Ping(ctx, &readpref.ReadPref{}); err != nil {
			log.Fatal(err)
		}
	}()

	var (
		port         = os.Getenv("PORT")
		dbname       = os.Getenv("DBNAME")
		userCollName = os.Getenv("USERCOLL")
		userColl     = client.Database(dbname).Collection(userCollName)
		userStore    = storage.NewMongoUserStore(client, userColl)
		store        = &storage.Store{User: userStore}
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
