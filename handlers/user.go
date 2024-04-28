package handlers

import (
	"fmt"
	"net/http"

	"github.com/HsiaoCz/search-engine/types"
	"github.com/gofiber/fiber/v2"
)

type UserHandlers struct {
}

func NewUserHandlers() *UserHandlers {
	return &UserHandlers{}
}

func (u *UserHandlers) HandleUserLogin(c *fiber.Ctx) error {
	var input types.Loginform
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "check out the input",
		})
	}
	fmt.Println(input)
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "login success!",
	})
}
