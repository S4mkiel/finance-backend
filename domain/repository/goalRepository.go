package repository

import (
	"context"
	"github.com/S4mkiel/finance-backend/domain/dto"
	"github.com/S4mkiel/finance-backend/domain/entity"
)

type GoalRepository interface {
	Get(context.Context, dto.GormQuery) ([]*entity.Goal, error)
	Find(context.Context, dto.GormQuery) (*entity.Goal, error)
	Create(context.Context, *entity.Goal) (*entity.Goal, error)
	Update(context.Context, *entity.Goal) (*entity.Goal, error)
	Delete(context.Context, dto.GormQuery) error
}
