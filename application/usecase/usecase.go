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

func (uc *UseCase) CreateUser(ctx context.Context, inDto *dto.CreateUserInDto) (*entity.User, error) {
	newUser, err := entity.NewUser(nil, inDto.Name, inDto.Password, inDto.Email)
	if err != nil {
		return nil, err
	}

	user, err := uc.userRepo.Create(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UseCase) FindUser(ctx context.Context, inDto *dto.FindUserInDto) (*string, error) {
	selectedUser, err := uc.userRepo.Find(ctx, queryDto.GormQuery{
		Where: &[]queryDto.GormWhere{
			{
				Column:    "email",
				Condition: "=",
				Value:     inDto.Email,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	if selectedUser == nil {
		return nil, errors.New("account not found")
	}

	isValid := utils.CompareHash(inDto.Password, selectedUser.Password)
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

func (uc *UseCase) GetUsers(ctx context.Context, inDto *dto.GetUsersInDto) ([]*entity.User, error) {
	users, err := uc.userRepo.Get(ctx, inDto.Query)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (uc *UseCase) CreateTransaction(ctx context.Context, inDto *dto.CreateTransactionInDto) (*entity.Transaction, error) {
	user, err := uc.userRepo.Find(ctx, queryDto.GormQuery{
		Where: &[]queryDto.GormWhere{
			{
				Column:    "id",
				Condition: "=",
				Value:     inDto.UserID,
			},
		},
	})

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("account not found")
	}

	transactionType, err := entity.NewTransactionType(*inDto.TransactionType)
	if err != nil {
		return nil, err
	}

	category, err := entity.NewTransactionCategory(*inDto.Category)
	if err != nil {
		return nil, err
	}

	newTransaction, err := entity.NewTransaction(nil, inDto.Amount, transactionType, category, inDto.Date, inDto.Notes, inDto.Currency, user)
	if err != nil {
		return nil, err
	}

	transaction, err := uc.transactionRepo.Create(ctx, newTransaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (uc *UseCase) GetTransactions(ctx context.Context, inDto *dto.GetTransactionsInDto) ([]*entity.Transaction, error) {
	user, err := uc.userRepo.Find(ctx, queryDto.GormQuery{
		Where: &[]queryDto.GormWhere{
			{
				Column:    "id",
				Condition: "=",
				Value:     inDto.UserID,
			},
		},
	})

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("account not found")
	}

	transactions, err := uc.transactionRepo.Get(ctx, queryDto.GormQuery{
		Where: &[]queryDto.GormWhere{
			{
				Column:    "user_id",
				Condition: "=",
				Value:     inDto.UserID,
			},
			{
				Column:    "amount",
				Condition: ">=",
				Value:     utils.PFloat64IfNotNil(inDto.StartAmount),
			},
			{
				Column:    "amount",
				Condition: "<=",
				Value:     utils.PFloat64IfNotNil(inDto.EndAmount),
			},
			{
				Column:    "transaction_type",
				Condition: "=",
				Value:     utils.PIntIfNotNil(inDto.TransactionType),
			},
			{
				Column:    "category",
				Condition: "=",
				Value:     utils.PIntIfNotNil(inDto.Category),
			},
			{
				Column:    "date",
				Condition: ">=",
				Value:     inDto.StartDate,
			},
			{
				Column:    "date",
				Condition: "<=",
				Value:     inDto.EndDate,
			},
			{
				Column:    "currency",
				Condition: "=",
				Value:     utils.PStringIfNotNil(inDto.Currency),
			},
			{
				Column:    "notes",
				Condition: "=",
				Value:     utils.PStringIfNotNil(inDto.Notes),
			},
		},
	})

	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (uc *UseCase) FindTransaction(ctx context.Context, inDto *dto.FindTransactionInDto) (*entity.Transaction, error) {
	user, err := uc.userRepo.Find(ctx, queryDto.GormQuery{
		Where: &[]queryDto.GormWhere{
			{
				Column:    "id",
				Condition: "=",
				Value:     inDto.UserID,
			},
		},
	})

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("account not found")
	}

	transaction, err := uc.transactionRepo.Find(ctx, queryDto.GormQuery{
		Where: &[]queryDto.GormWhere{
			{
				Column:    "id",
				Condition: "=",
				Value:     inDto.UserID,
			},
			{
				Column:    "user_id",
				Condition: "=",
				Value:     user.ID,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	if transaction == nil {
		return nil, errors.New("transaction not found")
	}

	return transaction, nil
}
