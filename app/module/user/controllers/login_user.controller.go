package user_controllers

import (
	"github.com/gofiber/fiber/v2"
	common_response "github.com/phucpham-infinity/go-nextpress/app/common/response"
	"github.com/phucpham-infinity/go-nextpress/app/context"
	user_model "github.com/phucpham-infinity/go-nextpress/app/module/user/model"
)

func (uc *userController) LoginUser(c *fiber.Ctx) error {

	payload := new(user_model.LoginUserParams)
	if err := c.BodyParser(payload); err != nil {
		return common_response.NewAppError(c).IsBadRequest(err).SendJSON()
	}
	validate := context.AppContext().GetValidator()
	if err := validate.Struct(payload); err != nil {
		return common_response.NewAppError(c).IsValidationError(err).SendJSON()
	}
	user, err := uc.userServices.LoginUser(c.Context(), payload)
	if err != nil {
		return common_response.NewAppError(c).IsBadRequest(err).SendJSON()
	}
	return common_response.NewSuccessResponse(c).WithData(user).SendJSON()

}
