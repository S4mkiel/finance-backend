package usecase

import (
	"context"
	"errors"
	"github.com/S4mkiel/finance-backend/application/dto"
	queryDto "github.com/S4mkiel/finance-backend/domain/dto"
	"github.com/S4mkiel/finance-backend/domain/entity"
	"github.com/S4mkiel/finance-backend/domain/repository"
	"github.com/S4mkiel/finance-backend/utils"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module(
	"usecase",
	fx.Provide(NewUsecase),
)

type UseCase struct {
	userRepo        repository.UserRepository
	transactionRepo repository.TransactionRepository
	rTransaction    repository.RecurringTransactionRepository
	budgetRepo      repository.BudgetRepository
	goalRepo        repository.GoalRepository
	logger          *zap.SugaredLogger
}

func NewUsecase(
	logger *zap.SugaredLogger,
	userRepo repository.UserRepository,
	transactionRepo repository.TransactionRepository,
	rTransactionRepo repository.RecurringTransactionRepository,
	budgetRepo repository.BudgetRepository,
	goalRepo repository.GoalRepository,
) (*UseCase, error) {
	return &UseCase{
		userRepo:        userRepo,
		transactionRepo: transactionRepo,
		rTransaction:    rTransactionRepo,
		budgetRepo:      budgetRepo,
		goalRepo:        goalRepo,
		logger:          logger,
	}, nil
}

func (uc *UseCase) CreateUser(ctx context.Context, dto *dto.CreateUserInDto) (*entity.User, error) {
	newUser, err := entity.NewUser(nil, dto.Name, dto.Password, dto.Email)
	if err != nil {
		return nil, err
	}

	user, err := uc.userRepo.Create(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UseCase) Find(ctx context.Context, dto *dto.FindUserInDto) (*string, error) {
	selectedUser, err := uc.userRepo.Find(ctx, queryDto.GormQuery{
		Where: &[]queryDto.GormWhere{
			{
				Column:    "email",
				Condition: "=",
				Value:     dto.Email,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	if selectedUser == nil {
		return nil, errors.New("account not found")
	}

	isValid := utils.CompareHash(dto.Password, selectedUser.Password)
	if !isValid {
		return nil, errors.New("invalid password")
	} else {
		token, err := utils.GenerateJWT(selectedUser.ID)
		if err != nil {
			return nil, err
		}
		return token, nil
	}
}

func (uc *UseCase) GetUsers(ctx context.Context, dto *dto.GetUsersInDto) ([]*entity.User, error) {
	users, err := uc.userRepo.Get(ctx, dto.Query)
	if err != nil {
		return nil, err
	}

	return users, nil
}
