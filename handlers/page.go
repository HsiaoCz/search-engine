package handlers

import (
	"net/http"

	"github.com/HsiaoCz/search-engine/types"
	"github.com/HsiaoCz/search-engine/views"
	"github.com/gofiber/fiber/v2"
)

type PageHandler struct{}

func NewPageHandler() *PageHandler {
	return &PageHandler{}
}

func (p *PageHandler) HandleHome(c *fiber.Ctx) error {
	return Render(c, views.Home())
}

func (p *PageHandler) HandleCrash(c *fiber.Ctx) error {
	var input types.SettingsForm
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "please check the input",
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": http.StatusOK,
		"data":   input,
	})
}
