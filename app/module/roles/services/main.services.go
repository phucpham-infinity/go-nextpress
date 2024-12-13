package roles_services

import (
	"context"

	"github.com/google/uuid"
	roles_database "github.com/phucpham-infinity/go-nextpress/app/module/roles/database"
	roles_model "github.com/phucpham-infinity/go-nextpress/app/module/roles/model"
)

type RolesServices interface {
	GetAllRoles(ctx context.Context) ([]roles_database.Role, error)
	GetRoles(ctx context.Context, params *roles_database.GetManyRolesParams) ([]roles_database.Role, error)
	GetByIdRoles(ctx context.Context, id uuid.UUID) (roles_database.Role, error)
	UpdateRoles(ctx context.Context, data *roles_database.UpdateRoleByIdParams) (roles_database.Role, error)
	DeleteByIdRoles(ctx context.Context, id uuid.UUID) (roles_database.Role, error)
	CreateRoles(ctx context.Context, data *roles_model.CreateRolesBody) (roles_database.Role, error)
}

type rolesServices struct {
	storage *roles_database.Queries
}

func NewRolesServices(storage *roles_database.Queries) *rolesServices {
	return &rolesServices{storage: storage}
}
