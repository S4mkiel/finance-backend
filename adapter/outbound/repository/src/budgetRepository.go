package src

import (
	"context"
	"errors"
	postgres "github.com/S4mkiel/finance-backend/adapter/outbound/database"
	"github.com/S4mkiel/finance-backend/domain/dto"
	"github.com/S4mkiel/finance-backend/domain/entity"
	"github.com/S4mkiel/finance-backend/domain/repository"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var BudgetModule = fx.Module(
	"recurring_transaction_repository",
	fx.Provide(NewBudgetRepositorySrc),
	fx.Provide(func(rsc *BudgetRepositorySrc) repository.BudgetRepository { return rsc }),
)

type BudgetRepositorySrc struct {
	pg     *postgres.Postgres
	logger *zap.SugaredLogger
}

func NewBudgetRepositorySrc(
	pg *postgres.Postgres,
	logger *zap.SugaredLogger,
) (*BudgetRepositorySrc, error) {
	return &BudgetRepositorySrc{
		pg:     pg,
		logger: logger,
	}, nil
}

func (s *BudgetRepositorySrc) Find(ctx context.Context, query dto.GormQuery) (*entity.Budget, error) {
	var item entity.Budget
	gormDB := QueryConstructor(s.pg.Db, query)
	result := gormDB.WithContext(ctx).First(&item)
	if result.Error != nil {
		if errors.Is(gorm.ErrRecordNotFound, result.Error) {
			return nil, nil
		} else {
			return nil, result.Error
		}
	}
	return &item, nil
}

func (s *BudgetRepositorySrc) Get(ctx context.Context, query dto.GormQuery) ([]*entity.Budget, error) {
	var item []*entity.Budget
	gormDB := QueryConstructor(s.pg.Db, query)
	result := gormDB.WithContext(ctx).Find(&item)
	if result.Error != nil {
		if errors.Is(gorm.ErrRecordNotFound, result.Error) {
			return item, nil
		} else {
			return nil, result.Error
		}
	}
	return item, nil
}

func (s *BudgetRepositorySrc) Create(ctx context.Context, goal *entity.Budget) (*entity.Budget, error) {
	if result := s.pg.Db.WithContext(ctx).Create(goal); result.Error != nil {
		return nil, result.Error
	}

	return goal, nil
}

func (s *BudgetRepositorySrc) Update(ctx context.Context, budget *entity.Budget) (*entity.Budget, error) {
	result := s.pg.Db.WithContext(ctx).Save(budget)
	if result.Error != nil {
		return nil, result.Error
	}

	return budget, nil
}

func (s *BudgetRepositorySrc) Delete(ctx context.Context, query dto.GormQuery) error {
	var item entity.Budget
	gormDB := QueryConstructor(s.pg.Db, query)
	result := gormDB.WithContext(ctx).Delete(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
