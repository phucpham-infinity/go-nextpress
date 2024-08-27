package user_controllers

import (
	"github.com/phucpham-infinity/go-nextpress/app/context"
	user_services "github.com/phucpham-infinity/go-nextpress/app/module/user/services"
	user_storage "github.com/phucpham-infinity/go-nextpress/app/module/user/storage"
)

type UserController interface {
	CreateUser() error
}
type userController struct {
	userServices user_services.UserServices
}

func NewUserController() *userController {
	db := context.AppContext().GetMySqlConnection()
	userStorage := user_storage.NewUserStorage(db)
	userServices := user_services.NewUserServices(userStorage)

	return &userController{
		userServices: userServices,
	}

}
