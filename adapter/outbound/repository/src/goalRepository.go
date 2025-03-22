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

var GoalModule = fx.Module(
	"recurring_transaction_repository",
	fx.Provide(NewGoalRepositorySrc),
	fx.Provide(func(rsc *GoalRepositorySrc) repository.GoalRepository { return rsc }),
)

type GoalRepositorySrc struct {
	pg     *postgres.Postgres
	logger *zap.SugaredLogger
}

func NewGoalRepositorySrc(
	pg *postgres.Postgres,
	logger *zap.SugaredLogger,
) (*GoalRepositorySrc, error) {
	return &GoalRepositorySrc{
		pg:     pg,
		logger: logger,
	}, nil
}

func (s *GoalRepositorySrc) Find(ctx context.Context, query dto.GormQuery) (*entity.Goal, error) {
	var item entity.Goal
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

func (s *GoalRepositorySrc) Get(ctx context.Context, query dto.GormQuery) ([]*entity.Goal, error) {
	var item []*entity.Goal
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

func (s *GoalRepositorySrc) Create(ctx context.Context, goal *entity.Goal) (*entity.Goal, error) {
	if result := s.pg.Db.WithContext(ctx).Create(goal); result.Error != nil {
		return nil, result.Error
	}

	return goal, nil
}

func (s *GoalRepositorySrc) Update(ctx context.Context, goal *entity.Goal) (*entity.Goal, error) {
	result := s.pg.Db.WithContext(ctx).Save(goal)
	if result.Error != nil {
		return nil, result.Error
	}

	return goal, nil
}

func (s *GoalRepositorySrc) Delete(ctx context.Context, query dto.GormQuery) error {
	var item entity.Goal
	gormDB := QueryConstructor(s.pg.Db, query)
	result := gormDB.WithContext(ctx).Delete(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
