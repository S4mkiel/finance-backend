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

var RecurringTransactionHandlerModule = fx.Module(
	"recurring_transaction_handler",
	fx.Provide(NewRecurringTransactionHandler),
)

type RecurringTransactionHandler struct {
	Uc     *usecase.UseCase
	logger *zap.SugaredLogger
}

func NewRecurringTransactionHandler(
	uc *usecase.UseCase,
	logger *zap.SugaredLogger,
) (*RecurringTransactionHandler, error) {
	return &RecurringTransactionHandler{Uc: uc, logger: logger}, nil
}

func (h *RecurringTransactionHandler) RegisterRoutes(r fiber.Router) {
	rTransaction := r.Group("/recurring-transaction", middleware.RateLimitMiddleware(), middleware.AuthMiddleware)
	rTransaction.Post("/", h.Create)
	rTransaction.Get("/", h.Get)
	rTransaction.Get("/:id", h.Find)
}

func (h *RecurringTransactionHandler) Create(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)

	var in dto.CreateRecurringTransactionInDto

	if err := c.BodyParser(&in); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.BaseDto{
			Success: utils.PBool(false),
			Error:   utils.PString(err.Error()),
		})
	}

	in.UserID = utils.PString(userID)

	rTransaction, httpResp, err := h.Uc.CreateRecurringTransaction(c.Context(), &in)
	if err != nil {
		return c.Status(*httpResp).JSON(dto.BaseDto{
			Success: utils.PBool(false),
			Error:   utils.PString(err.Error()),
		})
	}

	return c.Status(*httpResp).JSON(dto.BaseDto{
		Success: utils.PBool(true),
		Message: utils.PString("recurring transaction created"),
		Data:    rTransaction,
	})
}

func (h *RecurringTransactionHandler) Find(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	id := c.Params("id")

	in := &dto.FindRecurringTransactionInDto{
		ID:     utils.PString(id),
		UserID: utils.PString(userID),
	}

	rTransaction, httpResp, err := h.Uc.FindRecurringTransaction(c.Context(), in)
	if err != nil {
		return c.Status(*httpResp).JSON(dto.BaseDto{
			Success: utils.PBool(false),
			Error:   utils.PString(err.Error()),
		})
	}

	return c.Status(*httpResp).JSON(dto.BaseDto{
		Success: utils.PBool(true),
		Message: utils.PString("recurring transaction found"),
		Data:    rTransaction,
	})
}

func (h *RecurringTransactionHandler) Get(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)

	startAmount := c.QueryFloat("startAmount")
	endAmount := c.QueryFloat("endAmount")
	transactionType := c.QueryInt("type")
	category := c.QueryInt("category")
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")
	frequency := c.Query("frequency")
	currency := c.Query("currency")
	nextDate := c.Query("nextDate")

	in := dto.GetRecurringTransactionsInDto{
		UserID:          utils.PString(userID),
		StartAmount:     utils.Float64IfNotNil(startAmount),
		EndAmount:       utils.Float64IfNotNil(endAmount),
		TransactionType: utils.IntIfNotNil(transactionType),
		Category:        utils.IntIfNotNil(category),
		StartDate:       utils.TSTime(startDate),
		EndDate:         utils.TSTime(endDate),
		Frequency:       utils.StringIfNotNil(frequency),
		Currency:        utils.StringIfNotNil(currency),
		NextDate:        utils.TSTime(nextDate),
	}

	rTransaction, httpResp, err := h.Uc.GetRecurringTransaction(c.Context(), &in)
	if err != nil {
		return c.Status(*httpResp).JSON(dto.BaseDto{
			Success: utils.PBool(false),
			Error:   utils.PString(err.Error()),
		})
	}

	return c.Status(*httpResp).JSON(dto.BaseDto{
		Success: utils.PBool(true),
		Message: utils.PString("recurring transaction found"),
		Data:    rTransaction,
	})
}
