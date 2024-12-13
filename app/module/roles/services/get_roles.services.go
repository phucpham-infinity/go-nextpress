package roles_services

import (
	"context"

	roles_database "github.com/phucpham-infinity/go-nextpress/app/module/roles/database"
)

func (us *rolesServices) GetRoles(ctx context.Context, params *roles_database.GetManyRolesParams) ([]roles_database.Role, error) {
	return us.storage.GetManyRoles(ctx, *params)
}
