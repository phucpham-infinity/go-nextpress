// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.query.sql

package user_database

import (
	"context"
	"time"

	"github.com/google/uuid"
	null "gopkg.in/guregu/null.v4"
)

const activateUser = `-- name: ActivateUser :one
UPDATE users SET status = 'active', activation_key = '', updated_at = now(), registered_at = now() 
WHERE email = $1 
RETURNING id, username, email, activation_key, status, created_at, updated_at, registered_at, deleted_at
`

type ActivateUserRow struct {
	ID            uuid.NullUUID `json:"id"`
	Username      string        `json:"username"`
	Email         string        `json:"email"`
	ActivationKey string        `json:"activation_key"`
	Status        UserStatus    `json:"status"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
	RegisteredAt  null.Time     `json:"registered_at"`
	DeletedAt     null.Time     `json:"deleted_at"`
}

func (q *Queries) ActivateUser(ctx context.Context, email string) (ActivateUserRow, error) {
	row := q.db.QueryRowContext(ctx, activateUser, email)
	var i ActivateUserRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.ActivationKey,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.RegisteredAt,
		&i.DeletedAt,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (username, password, email, activation_key)
VALUES ($1, $2, $3, $4) 
RETURNING id, username, email, activation_key, status, created_at, updated_at, registered_at, deleted_at
`

type CreateUserParams struct {
	Username      string `json:"username"`
	Password      string `json:"password"`
	Email         string `json:"email"`
	ActivationKey string `json:"activation_key"`
}

type CreateUserRow struct {
	ID            uuid.NullUUID `json:"id"`
	Username      string        `json:"username"`
	Email         string        `json:"email"`
	ActivationKey string        `json:"activation_key"`
	Status        UserStatus    `json:"status"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
	RegisteredAt  null.Time     `json:"registered_at"`
	DeletedAt     null.Time     `json:"deleted_at"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.Password,
		arg.Email,
		arg.ActivationKey,
	)
	var i CreateUserRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.ActivationKey,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.RegisteredAt,
		&i.DeletedAt,
	)
	return i, err
}

const getManyUsers = `-- name: GetManyUsers :many
SELECT id, username, email, activation_key, status, created_at, updated_at, registered_at, deleted_at 
FROM users
`

type GetManyUsersRow struct {
	ID            uuid.NullUUID `json:"id"`
	Username      string        `json:"username"`
	Email         string        `json:"email"`
	ActivationKey string        `json:"activation_key"`
	Status        UserStatus    `json:"status"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
	RegisteredAt  null.Time     `json:"registered_at"`
	DeletedAt     null.Time     `json:"deleted_at"`
}

func (q *Queries) GetManyUsers(ctx context.Context) ([]GetManyUsersRow, error) {
	rows, err := q.db.QueryContext(ctx, getManyUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetManyUsersRow
	for rows.Next() {
		var i GetManyUsersRow
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Email,
			&i.ActivationKey,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.RegisteredAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, username, password, email, activation_key, status, created_at, updated_at, registered_at, deleted_at FROM users WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.ActivationKey,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.RegisteredAt,
		&i.DeletedAt,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, username, password, email, activation_key, status, created_at, updated_at, registered_at, deleted_at FROM users WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUserById(ctx context.Context, id uuid.NullUUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.ActivationKey,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.RegisteredAt,
		&i.DeletedAt,
	)
	return i, err
}
