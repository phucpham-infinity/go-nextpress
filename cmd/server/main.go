package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/phucpham-infinity/go-nextpress/app/configs"
	"github.com/phucpham-infinity/go-nextpress/app/global"
	"go.uber.org/zap"
)

func main() {
	configs.Run()

	app := fiber.New()

	app.Use(cors.New())
	app.Use(helmet.New())
	app.Use(requestid.New())
	app.Use(etag.New())
	app.Use(logger.New())
	app.Use(idempotency.New())
	app.Use(limiter.New())

	app.Get("/", func(c *fiber.Ctx) error {
		global.Logger.Info("Hello, World!", zap.String("okay", "success"))
		return c.SendString("Hello, World!")
	})

	app.Listen(global.Config.Server.Port)
}
