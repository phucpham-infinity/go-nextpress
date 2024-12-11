package user_router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/phucpham-infinity/go-nextpress/app/middlewares"
	user_controllers "github.com/phucpham-infinity/go-nextpress/app/module/user/controllers"
)

func NewUserRouter(v1 *fiber.Router) {
	userControllers := user_controllers.NewUserController()
	userRouter := (*v1).Group("/user")

	userRouter.Get("/", middlewares.AuthenticationMiddleware(), userControllers.GetManyUser)

	userRouter.Post("/register", userControllers.RegisterUser)
	userRouter.Post("/activate", userControllers.ActivateUser)
	userRouter.Get("/login", userControllers.LoginUser)
}
