package roles_services

import (
	"context"

	roles_database "github.com/phucpham-infinity/go-nextpress/app/module/roles/database"
	roles_model "github.com/phucpham-infinity/go-nextpress/app/module/roles/model"
)

func (us *rolesServices) CreateRoles(ctx context.Context, data *roles_model.CreateRolesBody) (roles_database.Role, error) {
	return us.storage.CreateRole(ctx, roles_database.CreateRoleParams{
		Role:       data.Role,
		Permission: data.Permission,
	})
}
