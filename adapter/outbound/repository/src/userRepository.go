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

var UserModule = fx.Module(
	"user_repository",
	fx.Provide(NewUserRepositorySrc),
	fx.Provide(func(rsc *UserRepositorySrc) repository.UserRepository { return rsc }),
)

type UserRepositorySrc struct {
	pg     *postgres.Postgres
	logger *zap.SugaredLogger
}

func NewUserRepositorySrc(
	pg *postgres.Postgres,
	logger *zap.SugaredLogger,
) (*UserRepositorySrc, error) {
	return &UserRepositorySrc{
		pg:     pg,
		logger: logger,
	}, nil
}

func (s *UserRepositorySrc) Find(ctx context.Context, query dto.GormQuery) (*entity.User, error) {
	var item entity.User
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

func (s *UserRepositorySrc) Get(ctx context.Context, query dto.GormQuery) ([]*entity.User, error) {
	var item []*entity.User
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

func (s *UserRepositorySrc) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	if result := s.pg.Db.WithContext(ctx).Create(user); result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (s *UserRepositorySrc) Update(ctx context.Context, user *entity.User) (*entity.User, error) {
	result := s.pg.Db.WithContext(ctx).Save(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (s *UserRepositorySrc) Delete(ctx context.Context, query dto.GormQuery) error {
	var item entity.User
	gormDB := QueryConstructor(s.pg.Db, query)
	result := gormDB.WithContext(ctx).Delete(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
