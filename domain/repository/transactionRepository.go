package repository

import (
	"context"
	"github.com/S4mkiel/finance-backend/domain/dto"
	"github.com/S4mkiel/finance-backend/domain/entity"
)

type TransactionRepository interface {
	Get(context.Context, dto.GormQuery) ([]*entity.Transaction, error)
	Find(context.Context, dto.GormQuery) (*entity.Transaction, error)
	Create(context.Context, *entity.Transaction) (*entity.Transaction, error)
	Update(context.Context, *entity.Transaction) (*entity.Transaction, error)
	Delete(context.Context, dto.GormQuery) error
}
