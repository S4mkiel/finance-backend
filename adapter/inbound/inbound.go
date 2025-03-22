package inbound

import (
	"github.com/S4mkiel/finance-backend/adapter/inbound/http"
	"go.uber.org/fx"
)

var Module = fx.Module("inbound", http.Module)
