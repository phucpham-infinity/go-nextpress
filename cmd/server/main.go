package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/phucpham-infinity/go-nextpress/app/configs"
	"github.com/phucpham-infinity/go-nextpress/app/context"
	user_router "github.com/phucpham-infinity/go-nextpress/app/module/user/router"
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

	config := context.AppContext().GetConfig()

	v1 := app.Group("/v1")
	user_router.NewUserRouter(&v1)

	log.Fatal(app.Listen(config.Server.Port))
}
