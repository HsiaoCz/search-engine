package routers

import (
	"github.com/HsiaoCz/search-engine/handlers"
	"github.com/HsiaoCz/search-engine/storage"
	"github.com/gofiber/fiber/v2"
)

func SetRouters(app *fiber.App, store *storage.Store) {
	var (
		pageHandler = handlers.NewPageHandler()
		userHandelr = handlers.NewUserHandlers(store)
	)
	// page handlers
	app.Get("/", pageHandler.HandleHome)
	app.Post("/", pageHandler.HandleCrash)

	// user handlers
	app.Post("/login", userHandelr.HandleUserLogin)
}
