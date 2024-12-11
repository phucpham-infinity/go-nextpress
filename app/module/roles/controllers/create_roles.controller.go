package roles_controllers

import (
	"github.com/gofiber/fiber/v2"
	common_response "github.com/phucpham-infinity/go-nextpress/app/common/response"
	"github.com/phucpham-infinity/go-nextpress/app/context"
	roles_model "github.com/phucpham-infinity/go-nextpress/app/module/roles/model"
)

func (uc *rolesController) CreateRoles(c *fiber.Ctx) error {

	payload := new(roles_model.CreateRolesBody)
	if err := c.BodyParser(payload); err != nil {
		return common_response.NewAppError(c).IsBadRequest(err).SendJSON()
	}
	validate := context.AppContext().GetValidator()
	if err := validate.Struct(payload); err != nil {
		return common_response.NewAppError(c).IsValidationError(err).SendJSON()
	}

	user, err := uc.userRolesServices.CreateRoles(c.Context(), payload)
	if err != nil {
		return common_response.NewAppError(c).IsBadRequest(err).SendJSON()
	}
	return common_response.NewSuccessResponse(c).WithData(user).SendJSON()

}