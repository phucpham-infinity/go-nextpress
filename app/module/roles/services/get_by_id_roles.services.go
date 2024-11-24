package roles_services

import (
	"context"

	"github.com/google/uuid"
	roles_database "github.com/phucpham-infinity/go-nextpress/app/module/roles/database"
)

func (us *rolesServices) GetByIdRoles(ctx context.Context, id uuid.UUID) (roles_database.Role, error) {
	return us.storage.GetByIdRoles(ctx, id)
}
