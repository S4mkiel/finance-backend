package repository

import (
	"context"
	"github.com/S4mkiel/finance-backend/domain/dto"
	"github.com/S4mkiel/finance-backend/domain/entity"
)

type RecurringTransactionRepository interface {
	Get(context.Context, dto.GormQuery) ([]*entity.RecurringTransaction, error)
	Find(context.Context, dto.GormQuery) (*entity.RecurringTransaction, error)
	Create(context.Context, *entity.RecurringTransaction) (*entity.RecurringTransaction, error)
	Update(context.Context, *entity.RecurringTransaction) (*entity.RecurringTransaction, error)
	Delete(context.Context, dto.GormQuery) error
}
