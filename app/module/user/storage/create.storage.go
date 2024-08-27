package user_storage

import (
	"context"

	user_model "github.com/phucpham-infinity/go-nextpress/app/module/user/model"
)

func (s *userStore) CreateUser(ctx context.Context, data *user_model.UserCreateStorage) error {
	if err := s.db.Table("users").Create(data).Error; err != nil {
		return err
	}
	return nil
}
