package application

import (
	"github.com/S4mkiel/finance-backend/application/usecase"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"application",
	usecase.Module,
)
