package user_services

import (
	"context"
	"errors"

	user_database "github.com/phucpham-infinity/go-nextpress/app/module/user/database"
	user_model "github.com/phucpham-infinity/go-nextpress/app/module/user/model"
)

func (us *userServices) ActivateUser(ctx context.Context, data *user_model.ActivateUserPrams) (user_database.ActivateUserRow, error) {
	email := data.Email

	user, err := us.storage.GetUserByEmail(ctx, email)
	if err != nil {
		return user_database.ActivateUserRow{}, err
	}

	if user.Status == user_database.UserStatusActive {
		return user_database.ActivateUserRow{}, errors.New("user is already activated")
	}

	return us.storage.ActivateUser(ctx, email)
}
