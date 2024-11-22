package user_services

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"time"

	user_database "github.com/phucpham-infinity/go-nextpress/app/module/user/database"
	user_model "github.com/phucpham-infinity/go-nextpress/app/module/user/model"
)

func (us *userServices) ActivateUser(ctx context.Context, data *user_model.ActivateUserParams) (user_database.ActivateUserRow, error) {
	email := data.Email

	user, err := us.storage.GetUserByEmail(ctx, email)
	if err != nil {
		return user_database.ActivateUserRow{}, err
	}

	if user.Status == user_database.UserStatusActive {
		return user_database.ActivateUserRow{}, errors.New("user is already activated")
	}

	storeKeyParts := strings.Split(user.ActivationKey, "::")
	if len(storeKeyParts) != 2 {
		return user_database.ActivateUserRow{}, errors.New("invalid activation key format")
	}

	storeKey := storeKeyParts[0]
	expirationTime, err := strconv.ParseInt(storeKeyParts[1], 10, 64)
	if err != nil {
		return user_database.ActivateUserRow{}, errors.New("invalid activation key format")
	}
	if time.Now().Unix() > expirationTime {
		return user_database.ActivateUserRow{}, errors.New("activation key has expired")
	}

	if storeKey != data.ActivationKey {
		return user_database.ActivateUserRow{}, errors.New("invalid activation key")
	}

	return us.storage.ActivateUser(ctx, email)
}
