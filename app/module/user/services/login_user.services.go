package user_services

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"time"

	common_helper "github.com/phucpham-infinity/go-nextpress/app/common/helper"
	user_database "github.com/phucpham-infinity/go-nextpress/app/module/user/database"
	user_model "github.com/phucpham-infinity/go-nextpress/app/module/user/model"
	"golang.org/x/crypto/bcrypt"
)

func (us *userServices) LoginUser(ctx context.Context, data *user_model.LoginUserParams) (user_model.LoginPayload, error) {
	email := data.Email

	user, err := us.storage.GetUserByEmail(ctx, email)
	if err != nil {
		return user_model.LoginPayload{}, err
	}

	if user.Status == user_database.UserStatusInactive {
		return user_model.LoginPayload{}, errors.New("user is not activated")
	}

	storePasswordParts := strings.Split(user.Password, "::")
	if len(storePasswordParts) != 2 {
		return user_model.LoginPayload{}, errors.New("invalid activation key format")
	}

	storePassword := storePasswordParts[0]
	expirationTime, err := strconv.ParseInt(storePasswordParts[1], 10, 64)
	if err != nil {
		return user_model.LoginPayload{}, errors.New("invalid activation key format")
	}
	if time.Now().Unix() > expirationTime {
		return user_model.LoginPayload{}, errors.New("activation key has expired")
	}

	err = bcrypt.CompareHashAndPassword([]byte(storePassword), []byte(data.Password))
	if err != nil {
		return user_model.LoginPayload{}, errors.New("invalid email or password")
	}

	token, err := common_helper.GenerateJWT(user.ID.UUID.String(), user.Email, user.Username)
	if err != nil {
		return user_model.LoginPayload{}, errors.New("failed to generate token")
	}

	return user_model.LoginPayload{
		Token: token,
		User: user_database.CreateUserRow{
			ID:           user.ID,
			Email:        user.Email,
			Username:     user.Username,
			Status:       user.Status,
			RegisteredAt: user.RegisteredAt,
			CreatedAt:    user.CreatedAt,
			UpdatedAt:    user.UpdatedAt,
			DeletedAt:    user.DeletedAt,
		},
	}, nil
}
