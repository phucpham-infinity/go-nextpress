package roles_controllers

import (
	"github.com/gofiber/fiber/v2"
	common_response "github.com/phucpham-infinity/go-nextpress/app/common/response"
)

func (uc *rolesController) GetAllRoles(c *fiber.Ctx) error {
	user, err := uc.userRolesServices.GetAllRoles(c.Context())
	if err != nil {
		return common_response.NewAppError(c).IsBadRequest(err).SendJSON()
	}
	return common_response.NewSuccessResponse(c).WithData(user).SendJSON()
}
