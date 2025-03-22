package repository

import (
	"context"
	"github.com/S4mkiel/finance-backend/domain/dto"
	"github.com/S4mkiel/finance-backend/domain/entity"
)

type BudgetRepository interface {
	Get(context.Context, dto.GormQuery) ([]*entity.Budget, error)
	Find(context.Context, dto.GormQuery) (*entity.Budget, error)
	Create(context.Context, *entity.Budget) (*entity.Budget, error)
	Update(context.Context, *entity.Budget) (*entity.Budget, error)
	Delete(context.Context, dto.GormQuery) error
}
