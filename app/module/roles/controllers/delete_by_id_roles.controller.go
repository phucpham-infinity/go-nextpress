package roles_controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	common_response "github.com/phucpham-infinity/go-nextpress/app/common/response"
)

func (uc *rolesController) DeleteRolesById(c *fiber.Ctx) error {
	id := c.Params("id")
	idUUID, err := uuid.Parse(id)
	if err != nil {
		return common_response.NewAppError(c).IsBadRequest(err).SendJSON()
	}
	user, err := uc.userRolesServices.DeleteByIdRoles(c.Context(), idUUID)
	if err != nil {
		return common_response.NewAppError(c).IsBadRequest(err).SendJSON()
	}
	return common_response.NewSuccessResponse(c).WithData(user).SendJSON()
}
