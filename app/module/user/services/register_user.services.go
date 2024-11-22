package user_services

import (
	"context"
	"fmt"
	"time"

	common_helper "github.com/phucpham-infinity/go-nextpress/app/common/helper"
	user_database "github.com/phucpham-infinity/go-nextpress/app/module/user/database"
	user_model "github.com/phucpham-infinity/go-nextpress/app/module/user/model"
	"golang.org/x/crypto/bcrypt"
)

func (us *userServices) RegisterUser(ctx context.Context, data *user_model.UserRegisterParams) (user_database.CreateUserRow, error) {
	email := data.Email
	username := data.Username

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return user_database.CreateUserRow{}, err
	}
	expirationTime := time.Now().Add(90 * 24 * time.Hour).Unix()
	finalPassword := fmt.Sprintf("%s::%d", string(hashedPassword), expirationTime)

	activationKey, err := common_helper.GenerateActivationKey(32)
	if err != nil {
		return user_database.CreateUserRow{}, err
	}
	activationExpirationTime := time.Now().Add(24 * time.Hour).Unix()
	finalActivationKey := fmt.Sprintf("%s::%d", activationKey, activationExpirationTime)

	return us.storage.CreateUser(ctx, user_database.CreateUserParams{
		Email:         email,
		Password:      finalPassword,
		Username:      username,
		ActivationKey: finalActivationKey,
	})
}
