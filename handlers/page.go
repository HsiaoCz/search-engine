package handlers

import (
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

func (p *PageHandler) HandleLogin(c *fiber.Ctx) error {
	return Render(c, views.Login())
}
