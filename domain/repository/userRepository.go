package repository

import (
	"context"
	"github.com/S4mkiel/finance-backend/domain/dto"
	"github.com/S4mkiel/finance-backend/domain/entity"
)

type UserRepository interface {
	Get(context.Context, dto.GormQuery) ([]*entity.User, error)
	Find(context.Context, dto.GormQuery) (*entity.User, error)
	Create(context.Context, *entity.User) (*entity.User, error)
	Update(context.Context, *entity.User) (*entity.User, error)
	Delete(context.Context, dto.GormQuery) error
}
