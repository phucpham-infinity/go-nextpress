package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/phucpham-infinity/go-nextpress/app/configs"
)

func main() {
	configs.Run()

	app := fiber.New()
	app.Use(etag.New())

	app.Listen(":3000")
}
