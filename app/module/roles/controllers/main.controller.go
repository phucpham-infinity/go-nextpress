package roles_controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/phucpham-infinity/go-nextpress/app/context"

	roles_database "github.com/phucpham-infinity/go-nextpress/app/module/roles/database"
	roles_services "github.com/phucpham-infinity/go-nextpress/app/module/roles/services"
)

type RolesController interface {
	GetAllRoles(c *fiber.Ctx) error
	GetByIdRoles(c *fiber.Ctx) error
}
type rolesController struct {
	userRolesServices roles_services.RolesServices
}

func NewRolesController() *rolesController {
	db := context.AppContext().GetMySqlConnection()

	rolesStorage := roles_database.New(db)

	userRolesServices := roles_services.NewRolesServices(rolesStorage)

	return &rolesController{
		userRolesServices: userRolesServices,
	}

}
