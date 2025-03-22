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

var RecurringTransactionModule = fx.Module(
	"recurring_transaction_repository",
	fx.Provide(NewRecurringTransactionRepositorySrc),
	fx.Provide(func(rsc *RecurringTransactionRepositorySrc) repository.RecurringTransactionRepository { return rsc }),
)

type RecurringTransactionRepositorySrc struct {
	pg     *postgres.Postgres
	logger *zap.SugaredLogger
}

func NewRecurringTransactionRepositorySrc(
	pg *postgres.Postgres,
	logger *zap.SugaredLogger,
) (*RecurringTransactionRepositorySrc, error) {
	return &RecurringTransactionRepositorySrc{
		pg:     pg,
		logger: logger,
	}, nil
}

func (s *RecurringTransactionRepositorySrc) Find(ctx context.Context, query dto.GormQuery) (*entity.RecurringTransaction, error) {
	var item entity.RecurringTransaction
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

func (s *RecurringTransactionRepositorySrc) Get(ctx context.Context, query dto.GormQuery) ([]*entity.RecurringTransaction, error) {
	var item []*entity.RecurringTransaction
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

func (s *RecurringTransactionRepositorySrc) Create(ctx context.Context, rTransaction *entity.RecurringTransaction) (*entity.RecurringTransaction, error) {
	if result := s.pg.Db.WithContext(ctx).Create(rTransaction); result.Error != nil {
		return nil, result.Error
	}

	return rTransaction, nil
}

func (s *RecurringTransactionRepositorySrc) Update(ctx context.Context, rTransaction *entity.RecurringTransaction) (*entity.RecurringTransaction, error) {
	result := s.pg.Db.WithContext(ctx).Save(rTransaction)
	if result.Error != nil {
		return nil, result.Error
	}

	return rTransaction, nil
}

func (s *RecurringTransactionRepositorySrc) Delete(ctx context.Context, query dto.GormQuery) error {
	var item entity.RecurringTransaction
	gormDB := QueryConstructor(s.pg.Db, query)
	result := gormDB.WithContext(ctx).Delete(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
