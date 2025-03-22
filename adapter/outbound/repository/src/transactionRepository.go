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

var TransactionModule = fx.Module(
	"transaction_repository",
	fx.Provide(NewTransactionRepositorySrc),
	fx.Provide(func(rsc *TransactionRepositorySrc) repository.TransactionRepository { return rsc }),
)

type TransactionRepositorySrc struct {
	pg     *postgres.Postgres
	logger *zap.SugaredLogger
}

func NewTransactionRepositorySrc(
	pg *postgres.Postgres,
	logger *zap.SugaredLogger,
) (*TransactionRepositorySrc, error) {
	return &TransactionRepositorySrc{
		pg:     pg,
		logger: logger,
	}, nil
}

func (s *TransactionRepositorySrc) Find(ctx context.Context, query dto.GormQuery) (*entity.Transaction, error) {
	var item entity.Transaction
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

func (s *TransactionRepositorySrc) Get(ctx context.Context, query dto.GormQuery) ([]*entity.Transaction, error) {
	var item []*entity.Transaction
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

func (s *TransactionRepositorySrc) Create(ctx context.Context, transaction *entity.Transaction) (*entity.Transaction, error) {
	if result := s.pg.Db.WithContext(ctx).Create(transaction); result.Error != nil {
		return nil, result.Error
	}

	return transaction, nil
}

func (s *TransactionRepositorySrc) Update(ctx context.Context, transaction *entity.Transaction) (*entity.Transaction, error) {
	result := s.pg.Db.WithContext(ctx).Save(transaction)
	if result.Error != nil {
		return nil, result.Error
	}

	return transaction, nil
}

func (s *TransactionRepositorySrc) Delete(ctx context.Context, query dto.GormQuery) error {
	var item entity.User
	gormDB := QueryConstructor(s.pg.Db, query)
	result := gormDB.WithContext(ctx).Delete(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
