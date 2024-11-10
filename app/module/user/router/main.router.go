package user_router

import (
	"github.com/gofiber/fiber/v2"
	user_controllers "github.com/phucpham-infinity/go-nextpress/app/module/user/controllers"
)

func NewUserRouter(v1 *fiber.Router) {
	userControllers := user_controllers.NewUserController()
	userUserRouter := (*v1).Group("/user")

	userUserRouter.Get("/", userControllers.GetManyUser)
	userUserRouter.Post("/register", userControllers.RegisterUser)
}
