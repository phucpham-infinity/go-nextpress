package user_services

import (
	"context"

	user_database "github.com/phucpham-infinity/go-nextpress/app/module/user/database"
)

func (us *userServices) GetManyUser(ctx context.Context) ([]user_database.GetManyUsersRow, error) {
	return us.storage.GetManyUsers(ctx)
}
