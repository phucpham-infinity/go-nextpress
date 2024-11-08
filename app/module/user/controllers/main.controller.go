package user_controllers

import (
	"github.com/phucpham-infinity/go-nextpress/app/context"
	user_database "github.com/phucpham-infinity/go-nextpress/app/module/user/database"
	user_services "github.com/phucpham-infinity/go-nextpress/app/module/user/services"
)

type UserController interface {
	RegisterUser() error
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
