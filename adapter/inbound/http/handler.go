package http

import (
	"github.com/S4mkiel/finance-backend/adapter/inbound/http/handler"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

var HandlerModule = fx.Module(
	"handler",
	handler.SwaggerHandlerModule,
	handler.UserHandlerModule,
	handler.UserNoAuthHandlerModule,
	fx.Invoke(HandleRoutes),
)

func HandleRoutes(
	http *Http,
	swaggerHandler *handler.SwaggerHandler,
	userHandler *handler.UserHandler,
	userNoAuthHandler *handler.UserNoAuthHandler,
) {
	http.App.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/v1/swagger/index.html")
	})

	v1 := http.App.Group("/v1")

	noAuth := http.App.Group("/v1/no-auth")

	swaggerHandler.RegisterRoutes(v1)
	userHandler.RegisterRoutes(v1)
	userNoAuthHandler.RegisterRoutes(noAuth)
}
