package roles_services

import (
	"context"
	roles_database "github.com/phucpham-infinity/go-nextpress/app/module/roles/database"
)

func (us *rolesServices) UpdateRoles(ctx context.Context, data *roles_database.UpdateRoleByIdParams) (roles_database.Role, error) {
	return us.storage.UpdateRoleById(ctx, *data)
}
