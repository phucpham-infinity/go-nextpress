package user_services

import (
	"context"

	user_model "github.com/phucpham-infinity/go-nextpress/app/module/user/model"
)

func (us *userServices) CreateUser(ctx context.Context, data *user_model.UserCreateStorage) error {
	return us.storage.CreateUser(ctx, data)
}
