package user_model

import (
	_ "github.com/go-playground/validator/v10"
	user_database "github.com/phucpham-infinity/go-nextpress/app/module/user/database"
)

type UserRegisterParams struct {
	Email         string `json:"email" validate:"required,email"`
	Username      string `json:"username" validate:"required,min=3,max=32"`
	Password      string `json:"password" validate:"required,min=4"`
	ActivationKey string
}

type ActivateUserParams struct {
	Email         string `json:"email" validate:"required,email"`
	ActivationKey string `json:"activationKey" validate:"required,min=4"`
}

type LoginUserParams struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=4"`
}

type LoginPayload struct {
	Token string                      `json:"token"`
	User  user_database.CreateUserRow `json:"user"`
}
