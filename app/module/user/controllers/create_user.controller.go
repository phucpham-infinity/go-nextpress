package user_controllers

import (
	"github.com/gofiber/fiber/v2"
	common_response "github.com/phucpham-infinity/go-nextpress/app/common/response"
	user_model "github.com/phucpham-infinity/go-nextpress/app/module/user/model"
)

func CreateUser(uc *userController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := new(user_model.UserCreateStorage)

		if err := c.BodyParser(payload); err != nil {
			return common_response.NewAppError(c).IsBadRequest(err).SendJSON()
		}

		if err := uc.userServices.CreateUser(c.Context(), payload); err != nil {
			return common_response.NewAppError(c).IsBadRequest(err).SendJSON()
		}

		return common_response.NewSuccessResponse(c).WithData(payload).SendJSON()
	}
}
