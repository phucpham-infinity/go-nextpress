package roles_controllers

import (
	"github.com/gofiber/fiber/v2"
	common_response "github.com/phucpham-infinity/go-nextpress/app/common/response"
	roles_database "github.com/phucpham-infinity/go-nextpress/app/module/roles/database"
)

func (uc *rolesController) GetRoles(c *fiber.Ctx) error {
	var page = c.Query("page", "1")
	var limit = c.QueryInt("limit", 10)

	user, err := uc.userRolesServices.GetRoles(c.Context(), &roles_database.GetManyRolesParams{
		Page:      page,
		PageLimit: limit,
	})
	if err != nil {
		return common_response.NewAppError(c).IsBadRequest(err).SendJSON()
	}
	return common_response.NewSuccessResponse(c).WithData(user).SendJSON()
}
