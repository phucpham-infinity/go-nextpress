package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/phucpham-infinity/go-nextpress/pkg/response"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.GetRespHeader("Authorization")
		if token != "demo" {
			return response.ErrorResponse(c, response.StatusNetworkAuthenticationRequired)
		}
		return c.Next()
	}
}
