package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/phucpham-infinity/go-nextpress/pkg/response"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return response.SuccessResponse(c, 200, []string{"91"})
	})

	app.Listen(":3000")
}
