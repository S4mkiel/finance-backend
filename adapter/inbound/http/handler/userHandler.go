package handler

import (
	"github.com/S4mkiel/finance-backend/adapter/inbound/http/middleware"
	"github.com/S4mkiel/finance-backend/application/dto"
	"github.com/S4mkiel/finance-backend/application/usecase"
	"github.com/S4mkiel/finance-backend/utils"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var UserHandlerModule = fx.Module(
	"user_handler",
	fx.Provide(NewUserHandler),
)

type UserHandler struct {
	Uc     *usecase.UseCase
	logger *zap.SugaredLogger
}

func NewUserHandler(
	uc *usecase.UseCase,
	logger *zap.SugaredLogger,
) (*UserHandler, error) {
	return &UserHandler{Uc: uc, logger: logger}, nil
}

func (h *UserHandler) RegisterRoutes(r fiber.Router) {
	users := r.Group("/users", middleware.RateLimitMiddleware(), middleware.AuthMiddleware)
	users.Get("/", h.GetUsers)
}

// GetUsers godoc
// @Security ApiKeyAuth
// @Summary get users
// @ID getUsers
// @Tags User
// @Description Router for get all users
// @Accept json
// @Produce json
// @Success 200 {object} entity.User
// @Failure 400 {object} dto.BaseDto
// @Failure 403 {object} dto.BaseDto
// @Failure 404 {object} dto.BaseDto
// @Router /users/ [get]
func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	user, httpResp, err := h.Uc.GetUsers(c.Context(), &dto.GetUsersInDto{})
	if err != nil {
		return c.Status(*httpResp).JSON(dto.BaseDto{
			Success: utils.PBool(false),
			Error:   utils.PString(err.Error()),
		})
	}

	return c.Status(*httpResp).JSON(dto.BaseDto{
		Success: utils.PBool(true),
		Message: utils.PString("List of users"),
		Data:    user,
	})
}
