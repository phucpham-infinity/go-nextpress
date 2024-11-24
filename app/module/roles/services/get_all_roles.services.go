package roles_services

import (
	"context"

	roles_database "github.com/phucpham-infinity/go-nextpress/app/module/roles/database"
)

func (us *rolesServices) GetAllRoles(ctx context.Context) ([]roles_database.Role, error) {
	return us.storage.GetAllRoles(ctx)
}
