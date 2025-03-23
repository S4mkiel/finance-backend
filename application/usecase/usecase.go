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

func (uc *UseCase) CreateUser(ctx context.Context, inDto *dto.CreateUserInDto) (*entity.User, *int, error) {
	newUser, err := entity.NewUser(nil, inDto.Name, inDto.Password, inDto.Email)
	if err != nil {
		uc.logger.Error(err)
		return nil, utils.PInt(500), err
	}

	user, err := uc.userRepo.Create(ctx, newUser)
	if err != nil {
		uc.logger.Error(err)
		return nil, utils.PInt(500), err
	}

	return user, utils.PInt(200), nil
}

func (uc *UseCase) FindUser(ctx context.Context, inDto *dto.FindUserInDto) (*string, *int, error) {
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
		uc.logger.Error(err)
		return nil, utils.PInt(500), err
	}

	if selectedUser == nil {
		uc.logger.Warn("user not found")
		return nil, utils.PInt(404), errors.New("account not found")
	}

	isValid := utils.CompareHash(inDto.Password, selectedUser.Password)
	if !isValid {
		uc.logger.Warn("invalid password")
		return nil, utils.PInt(401), errors.New("invalid password")
	} else {
		token, err := utils.GenerateJWT(selectedUser.ID)
		if err != nil {
			uc.logger.Error(err)
			return nil, utils.PInt(500), err
		}
		return token, utils.PInt(200), nil
	}
}

func (uc *UseCase) GetUsers(ctx context.Context, inDto *dto.GetUsersInDto) ([]*entity.User, *int, error) {
	users, err := uc.userRepo.Get(ctx, inDto.Query)
	if err != nil {
		uc.logger.Error(err)
		return nil, utils.PInt(500), err
	}

	if users == nil || len(users) == 0 {
		uc.logger.Warn("users not found")
		return users, utils.PInt(404), errors.New("users not found")
	}

	return users, utils.PInt(200), nil
}

func (uc *UseCase) CreateTransaction(ctx context.Context, inDto *dto.CreateTransactionInDto) (*entity.Transaction, *int, error) {
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
		uc.logger.Error(err)
		return nil, utils.PInt(500), err
	}

	if user == nil {
		uc.logger.Warn("user not found")
		return nil, utils.PInt(404), errors.New("account not found")
	}

	transactionType, err := entity.NewTransactionType(*inDto.TransactionType)
	if err != nil {
		uc.logger.Error(err)
		return nil, utils.PInt(500), err
	}

	category, err := entity.NewTransactionCategory(*inDto.Category)
	if err != nil {
		uc.logger.Error(err)
		return nil, utils.PInt(500), err
	}

	newTransaction, err := entity.NewTransaction(nil, inDto.Amount, transactionType, category, inDto.Date, inDto.Notes, inDto.Currency, user)
	if err != nil {
		uc.logger.Error(err)
		return nil, utils.PInt(500), err
	}

	transaction, err := uc.transactionRepo.Create(ctx, newTransaction)
	if err != nil {
		uc.logger.Error(err)
		return nil, utils.PInt(500), err
	}

	return transaction, utils.PInt(201), nil
}

func (uc *UseCase) GetTransactions(ctx context.Context, inDto *dto.GetTransactionsInDto) ([]*entity.Transaction, *int, error) {
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
		uc.logger.Error(err)
		return nil, utils.PInt(500), err
	}

	if user == nil {
		uc.logger.Warn("user not found")
		return nil, utils.PInt(404), errors.New("account not found")
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
		uc.logger.Error(err)
		return nil, utils.PInt(500), err
	}

	if transactions == nil || len(transactions) == 0 {
		uc.logger.Warn("transaction not found")
		return transactions, utils.PInt(404), errors.New("transaction not found")
	}

	return transactions, utils.PInt(200), nil
}

func (uc *UseCase) FindTransaction(ctx context.Context, inDto *dto.FindTransactionInDto) (*entity.Transaction, *int, error) {
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
		uc.logger.Error(err)
		return nil, utils.PInt(500), err
	}

	if user == nil {
		uc.logger.Warn("user not found")
		return nil, utils.PInt(404), errors.New("account not found")
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
		uc.logger.Error(err)
		return nil, utils.PInt(500), err
	}

	if transaction == nil {
		uc.logger.Warn("transaction not found")
		return nil, utils.PInt(404), errors.New("transaction not found")
	}

	return transaction, utils.PInt(200), nil
}

func (uc *UseCase) CreateRecurringTransaction(ctx context.Context, inDto *dto.CreateRecurringTransactionInDto) (*entity.RecurringTransaction, *int, error) {
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
		uc.logger.Error(err)
		return nil, utils.PInt(500), err
	}

	if user == nil {
		uc.logger.Warn("user not found")
		return nil, utils.PInt(404), errors.New("account not found")
	}

	transactionType, err := entity.NewTransactionType(*inDto.TransactionType)
	if err != nil {
		uc.logger.Error(err)
		return nil, utils.PInt(500), err
	}

	category, err := entity.NewTransactionCategory(*inDto.Category)
	if err != nil {
		uc.logger.Error(err)
		return nil, utils.PInt(500), err
	}

	newRTransaction, err := entity.NewRecurringTransaction(nil, inDto.Amount, transactionType, category, inDto.Frequency, inDto.NextDate, user)
	if err != nil {
		uc.logger.Error(err)
		return nil, utils.PInt(500), err
	}

	rTransaction, err := uc.rTransaction.Create(ctx, newRTransaction)
	if err != nil {
		uc.logger.Error(err)
		return nil, utils.PInt(500), err
	}

	return rTransaction, utils.PInt(201), err
}

func (uc *UseCase) FindRecurringTransaction(ctx context.Context, inDto *dto.FindRecurringTransactionInDto) (*entity.RecurringTransaction, *int, error) {
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
		uc.logger.Error(err)
		return nil, utils.PInt(500), err
	}

	if user == nil {
		uc.logger.Warn("user not found")
		return nil, utils.PInt(404), errors.New("account not found")
	}

	rTransaction, err := uc.rTransaction.Find(ctx, queryDto.GormQuery{
		Where: &[]queryDto.GormWhere{
			{
				Column:    "id",
				Condition: "=",
				Value:     inDto.ID,
			},
			{
				Column:    "user_id",
				Condition: "=",
				Value:     user.ID,
			},
		},
	})
	if err != nil {
		uc.logger.Error(err)
		return nil, utils.PInt(500), err
	}

	if rTransaction == nil {
		uc.logger.Warn("transaction not found")
		return nil, utils.PInt(404), errors.New("transaction not found")
	}

	return rTransaction, utils.PInt(201), nil
}

func (uc *UseCase) GetRecurringTransaction(ctx context.Context, inDto *dto.GetRecurringTransactionsInDto) ([]*entity.RecurringTransaction, *int, error) {
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
		uc.logger.Error(err)
		return nil, utils.PInt(500), err
	}

	if user == nil {
		uc.logger.Warn("user not found")
		return nil, utils.PInt(404), errors.New("account not found")
	}

	rTransaction, err := uc.rTransaction.Get(ctx, queryDto.GormQuery{
		Where: &[]queryDto.GormWhere{
			{
				Column:    "user_id",
				Condition: "=",
				Value:     user.ID,
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
				Column:    "frequency",
				Condition: "=",
				Value:     utils.PStringIfNotNil(inDto.Frequency),
			},
			{
				Column:    "next_date",
				Condition: "=",
				Value:     inDto.NextDate,
			},
		},
	})

	if err != nil {
		uc.logger.Error(err)
		return nil, utils.PInt(500), err
	}

	if rTransaction == nil || len(rTransaction) == 0 {
		uc.logger.Warn("transaction not found")
		return rTransaction, utils.PInt(404), errors.New("transaction not found")
	}

	return rTransaction, utils.PInt(200), nil
}
