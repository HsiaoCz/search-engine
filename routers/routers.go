package routers

import (
	"github.com/HsiaoCz/search-engine/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetRouters(app *fiber.App) {
	var (
		pageHandler = handlers.NewPageHandler()
		userHandelr = handlers.NewUserHandlers()
	)
	// page handlers
	app.Get("/", pageHandler.HandleHome)
	app.Post("/", pageHandler.HandleCrash)

	// user handlers
	app.Post("/login", userHandelr.HandleUserLogin)
}
