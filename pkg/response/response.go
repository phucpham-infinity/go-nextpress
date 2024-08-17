package response

import (
	"github.com/gofiber/fiber/v2"
)

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(c *fiber.Ctx, code int, data interface{}) error {
	return c.Status(code).JSON(ResponseData{
		Code:    code,
		Message: StatusText(code),
		Data:    data,
	})
}

func ErrorResponse(c *fiber.Ctx, code int) error {
	return c.Status(code).JSON(ResponseData{
		Code:    code,
		Message: StatusText(code),
		Data:    nil,
	})
}
