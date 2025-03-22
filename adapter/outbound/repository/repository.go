package repository

import (
	"github.com/S4mkiel/finance-backend/adapter/outbound/repository/src"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"repository",
	src.UserModule,
	src.TransactionModule,
	src.BudgetModule,
	src.GoalModule,
	src.RecurringTransactionModule,
)
