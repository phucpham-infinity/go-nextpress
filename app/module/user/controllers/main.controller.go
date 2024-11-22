package user_controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/phucpham-infinity/go-nextpress/app/context"
	user_database "github.com/phucpham-infinity/go-nextpress/app/module/user/database"
	user_services "github.com/phucpham-infinity/go-nextpress/app/module/user/services"
)

type UserController interface {
	RegisterUser(c *fiber.Ctx) error
	ActivateUser(c *fiber.Ctx) error
	LoginUser(c *fiber.Ctx) error
	GetManyUser(c *fiber.Ctx) error
}
type userController struct {
	userServices user_services.UserServices
}

func NewUserController() *userController {
	db := context.AppContext().GetMySqlConnection()
	userStorage := user_database.New(db)
	userServices := user_services.NewUserServices(userStorage)

	return &userController{
		userServices: userServices,
	}

}
