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

var TransactionHandlerModule = fx.Module(
	"transaction_handler",
	fx.Provide(NewTransactionHandler),
)

type TransactionHandler struct {
	Uc     *usecase.UseCase
	logger *zap.SugaredLogger
}

func NewTransactionHandler(
	uc *usecase.UseCase,
	logger *zap.SugaredLogger,
) (*TransactionHandler, error) {
	return &TransactionHandler{Uc: uc, logger: logger}, nil
}

func (h *TransactionHandler) RegisterRoutes(r fiber.Router) {
	transaction := r.Group("/transaction", middleware.RateLimitMiddleware(), middleware.AuthMiddleware)
	transaction.Post("/", h.CreateTransaction)
	transaction.Get("/", h.GetTransactions)
	transaction.Get("/:id", h.FindTransaction)
}

// CreateTransaction godoc
// @Security ApiKeyAuth
// @Summary create transaction
// @ID createTransaction
// @Tags Transaction
// @Description Router for create transaction
// @Accept json
// @Produce json
// @Param _ body dto.CreateTransactionInDto true "create user payload"
// @Success 200 {object} entity.Transaction
// @Failure 400 {object} dto.BaseDto
// @Failure 403 {object} dto.BaseDto
// @Failure 404 {object} dto.BaseDto
// @Router /transaction/ [post]
func (h *TransactionHandler) CreateTransaction(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)

	var in dto.CreateTransactionInDto

	if err := c.BodyParser(&in); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.BaseDto{
			Success: utils.PBool(false),
			Error:   utils.PString(err.Error()),
		})
	}

	in.UserID = utils.PString(userID)

	transaction, err := h.Uc.CreateTransaction(c.Context(), &in)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.BaseDto{
			Success: utils.PBool(false),
			Error:   utils.PString(err.Error()),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(dto.BaseDto{
		Success: utils.PBool(true),
		Message: utils.PString("transaction created successfully"),
		Data:    transaction,
	})
}

// GetTransactions godoc
// @Security ApiKeyAuth
// @Summary Get transactions by user with filters
// @ID getTransactions
// @Tags Transaction
// @Description Retrieves transactions based on filters provided in query parameters
// @Accept json
// @Produce json
// @Param startAmount query float64 false "Minimum transaction amount"
// @Param endAmount query float64 false "Maximum transaction amount"
// @Param type query int false "Transaction type (0 for income, 1 for expense)"
// @Param category query int false "Transaction category"
// @Param startDate query string false "Start date (format: YYYY-MM-DDT00:00)"
// @Param endDate query string false "End date (format: YYYY-MM-DDT00:00)"
// @Param notes query string false "Search transactions by notes"
// @Param currency query string false "Filter transactions by currency"
// @Success 200 {object} entity.Transaction
// @Failure 400 {object} dto.BaseDto
// @Failure 403 {object} dto.BaseDto
// @Failure 404 {object} dto.BaseDto
// @Router /transaction/ [get]
func (h *TransactionHandler) GetTransactions(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)

	startAmount := c.Query("startAmount")
	endAmount := c.Query("endAmount")
	transactionType := c.QueryInt("type")
	category := c.QueryInt("category")
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")
	notes := c.Query("notes")
	currency := c.Query("currency")

	in := dto.GetTransactionsInDto{
		UserID:          utils.PString(userID),
		StartAmount:     utils.TSFloat64(startAmount),
		EndAmount:       utils.TSFloat64(endAmount),
		TransactionType: utils.IntIfNotNil(transactionType),
		Category:        utils.IntIfNotNil(category),
		StartDate:       utils.TSTime(startDate),
		EndDate:         utils.TSTime(endDate),
		Notes:           utils.StringIfNotNil(notes),
		Currency:        utils.StringIfNotNil(currency),
	}

	transaction, err := h.Uc.GetTransactions(c.Context(), &in)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.BaseDto{
			Success: utils.PBool(false),
			Error:   utils.PString(err.Error()),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(dto.BaseDto{
		Success: utils.PBool(true),
		Message: utils.PString("transactions founded with successfully"),
		Data:    transaction,
	})
}

// FindTransaction godoc
// @Security ApiKeyAuth
// @Summary find transaction by id
// @ID findTransaction
// @Tags Transaction
// @Description Router for find transaction by id
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} entity.Transaction
// @Failure 400 {object} dto.BaseDto
// @Failure 403 {object} dto.BaseDto
// @Failure 404 {object} dto.BaseDto
// @Router /transaction/{id} [get]
func (h *TransactionHandler) FindTransaction(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	id := c.Params("id")

	var in dto.FindTransactionInDto

	in.UserID = utils.PString(userID)
	in.ID = utils.PString(id)

	transaction, err := h.Uc.FindTransaction(c.Context(), &in)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.BaseDto{
			Success: utils.PBool(false),
			Error:   utils.PString(err.Error()),
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.BaseDto{
		Success: utils.PBool(true),
		Message: utils.PString("transaction found successfully"),
		Data:    transaction,
	})
}
