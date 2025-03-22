package dto

import (
	"github.com/S4mkiel/finance-backend/domain/dto"
	"time"
)

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

type CreateTransactionInDto struct {
	UserID          *string
	Amount          *float64
	TransactionType *int
	Category        *int
	Date            *time.Time
	Notes           *string
	Currency        *string
}

type GetTransactionsInDto struct {
	UserID          *string
	StartAmount     *float64
	EndAmount       *float64
	TransactionType *int
	Category        *int
	StartDate       *time.Time
	EndDate         *time.Time
	Notes           *string
	Currency        *string
}

type FindTransactionInDto struct {
	UserID *string
	ID     *string
}
