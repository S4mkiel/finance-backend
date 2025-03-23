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

var UserNoAuthHandlerModule = fx.Module(
	"user_no_auth_handler",
	fx.Provide(NewUserNoAuthHandler),
)

type UserNoAuthHandler struct {
	Uc     *usecase.UseCase
	logger *zap.SugaredLogger
}

func NewUserNoAuthHandler(
	uc *usecase.UseCase,
	logger *zap.SugaredLogger,
) (*UserNoAuthHandler, error) {
	return &UserNoAuthHandler{Uc: uc, logger: logger}, nil
}

func (h *UserNoAuthHandler) RegisterRoutes(r fiber.Router) {
	users := r.Group("/users", middleware.RateLimitMiddleware())
	users.Post("/signing", h.Signing)
	users.Post("/signup", h.CreateUser)
}

// CreateUser godoc
// @Summary create user
// @ID createUser
// @Tags User
// @Description Router for create user
// @Accept json
// @Produce json
// @Param _ body dto.CreateUserInDto true "create user payload"
// @Success 200 {object} entity.User
// @Failure 400 {object} dto.BaseDto
// @Failure 403 {object} dto.BaseDto
// @Failure 404 {object} dto.BaseDto
// @Router /no-auth/users/signup [post]
func (h *UserNoAuthHandler) CreateUser(c *fiber.Ctx) error {
	var in dto.CreateUserInDto

	if err := c.BodyParser(&in); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.BaseDto{
			Success: utils.PBool(false),
			Error:   utils.PString(err.Error()),
		})
	}

	user, httpResp, err := h.Uc.CreateUser(c.Context(), &in)
	if err != nil {
		return c.Status(*httpResp).JSON(dto.BaseDto{
			Success: utils.PBool(false),
			Error:   utils.PString(err.Error()),
		})
	}

	return c.Status(*httpResp).JSON(dto.BaseDto{
		Success: utils.PBool(true),
		Message: utils.PString("User created"),
		Data:    user,
	})
}

// Signing godoc
// @Summary signing
// @ID Signing
// @Tags User
// @Description Router for signing
// @Accept json
// @Produce json
// @Param _ body dto.FindUserInDto true "signing payload"
// @Failure 400 {object} dto.BaseDto
// @Failure 403 {object} dto.BaseDto
// @Failure 404 {object} dto.BaseDto
// @Router /no-auth/users/signing [post]
func (h *UserNoAuthHandler) Signing(c *fiber.Ctx) error {
	var in dto.FindUserInDto

	if err := c.BodyParser(&in); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.BaseDto{
			Success: utils.PBool(false),
			Error:   utils.PString(err.Error()),
		})
	}

	token, httpResp, err := h.Uc.FindUser(c.Context(), &in)
	if err != nil {
		return c.Status(*httpResp).JSON(dto.BaseDto{
			Success: utils.PBool(false),
			Error:   utils.PString(err.Error()),
		})
	}

	return c.Status(*httpResp).JSON(dto.BaseDto{
		Success: utils.PBool(true),
		Message: utils.PString("Sign in successful"),
		Data:    token,
	})
}
