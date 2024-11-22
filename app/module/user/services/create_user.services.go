package user_services

import (
	"context"

	user_database "github.com/phucpham-infinity/go-nextpress/app/module/user/database"
	user_model "github.com/phucpham-infinity/go-nextpress/app/module/user/model"
)

func (us *userServices) RegisterUser(ctx context.Context, data *user_model.UserRegisterStorage) (user_database.CreateUserRow, error) {
	email := data.Email
	password := data.Password
	username := data.Username
	activationKey := data.ActivationKey

	return us.storage.CreateUser(ctx, user_database.CreateUserParams{
		Email:         email,
		Password:      password,
		Username:      username,
		ActivationKey: activationKey,
	})
}
