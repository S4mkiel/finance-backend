package adapter

import (
	"github.com/S4mkiel/finance-backend/adapter/inbound"
	"github.com/S4mkiel/finance-backend/adapter/outbound"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"adapter",
	outbound.Module,
	inbound.Module,
)
