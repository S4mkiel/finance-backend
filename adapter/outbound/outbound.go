package outbound

import (
	postgres "github.com/S4mkiel/finance-backend/adapter/outbound/database"
	"github.com/S4mkiel/finance-backend/adapter/outbound/logger"
	"github.com/S4mkiel/finance-backend/adapter/outbound/repository"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"outbound",
	logger.Module,
	postgres.Module,
	repository.Module,
)
