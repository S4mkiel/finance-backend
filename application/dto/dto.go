package dto

import "github.com/S4mkiel/finance-backend/domain/dto"

type BaseDto struct {
	Success *bool   `json:"success;default:false"`
	Error   *string `json:"error,omitempty"`
	Message *string `json:"message,omitempty"`
	Data    any     `json:"data"`
}

type CreateUserInDto struct {
	Email    *string
	Name     *string
	Password *string
}
type GetUsersInDto struct {
	Query dto.GormQuery
}

type FindUserInDto struct {
	Email    *string
	Password *string
}
