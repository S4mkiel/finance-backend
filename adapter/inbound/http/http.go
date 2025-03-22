package http

import (
	"context"
	"fmt"
	_ "github.com/S4mkiel/finance-backend/adapter/inbound/http/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module(
	"http",
	ConfigModule,
	HandlerModule,
	fx.Provide(NewHttp),
	fx.Invoke(HookHttp),
)

// @title Finance api
// @version 1.0
// @description Swagger API for finance service.
// @termsOfService http://swagger.io/terms/

// @contact.name Roberto Filho
// @contact.email zferraz.rf@gmail.com

// @BasePath /v1
// @query.collection.format multi

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

type Http struct {
	App *fiber.App
}

func NewHttp(c *Config) (*Http, error) {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: *c.DisableStartupMessage,
	})
	app.Use(cors.New())

	return &Http{
		App: app,
	}, nil
}

func HookHttp(lc fx.Lifecycle, http *Http, l *zap.SugaredLogger, c *Config) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				l.Infof("listening http service on port %d", *c.Port)
				if err := http.App.Listen(fmt.Sprintf(":%d", *c.Port)); err != nil {
					l.Error(err)
					panic(err)
				}
			}()

			return nil
		},
		OnStop: func(context.Context) error {
			return http.App.Shutdown()
		},
	})
}
