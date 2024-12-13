package user_roles_router

import (
	"github.com/gofiber/fiber/v2"
	roles_controllers "github.com/phucpham-infinity/go-nextpress/app/module/roles/controllers"
)

func NewUserRolesRouter(v1 *fiber.Router) {
	rolesController := roles_controllers.NewRolesController()
	rolesRouter := (*v1).Group("/roles")

	//rolesRouter.Get("/", middlewares.AuthenticationMiddleware(), rolesController.GetAllRoles)
	//rolesRouter.Get("/:id", middlewares.AuthenticationMiddleware(), rolesController.GetByIdRoles)
	//
	//rolesRouter.Post("/", middlewares.AuthenticationMiddleware(), rolesController.CreateRoles)
	//rolesRouter.Put("/", middlewares.AuthenticationMiddleware(), rolesController.UpdateRoles)
	//
	//rolesRouter.Delete("/:id", middlewares.AuthenticationMiddleware(), rolesController.DeleteRolesById)

	rolesRouter.Get("/", rolesController.GetRoles)
	rolesRouter.Get("/all", rolesController.GetAllRoles)
	rolesRouter.Get("/:id", rolesController.GetByIdRoles)

	rolesRouter.Post("/", rolesController.CreateRoles)
	rolesRouter.Put("/", rolesController.UpdateRoles)

	rolesRouter.Delete("/:id", rolesController.DeleteRolesById)
}
