package __model__router

import (
	"github.com/gofiber/fiber/v2"
	middlewares "github.com/phucpham-infinity/go-nextpress/app/middlewares"
	user_controllers "github.com/phucpham-infinity/go-nextpress/app/module/user/controllers"
)

func NewUserRouter(v1 *fiber.Router) {
	userControllers := user_controllers.NewUserController()
	userUserRouter := (*v1).Group("/user")

	userUserRouter.Get("/", middlewares.AuthenticationMiddleware(), userControllers.GetManyUser)

	userUserRouter.Post("/register", userControllers.RegisterUser)
	userUserRouter.Post("/activate", userControllers.ActivateUser)
	userUserRouter.Get("/login", userControllers.LoginUser)
}
