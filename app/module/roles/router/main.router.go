package user_roles_router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/phucpham-infinity/go-nextpress/app/middlewares"
	roles_controllers "github.com/phucpham-infinity/go-nextpress/app/module/roles/controllers"
)

func NewUserRolesRouter(v1 *fiber.Router) {
	rolesController := roles_controllers.NewRolesController()
	rolesRouter := (*v1).Group("/roles")

	rolesRouter.Get("/", middlewares.AuthenticationMiddleware(), rolesController.GetAllRoles)
	rolesRouter.Get("/:id", middlewares.AuthenticationMiddleware(), rolesController.GetByIdRoles)
	rolesRouter.Post("/", middlewares.AuthenticationMiddleware(), rolesController.CreateRoles)

}
