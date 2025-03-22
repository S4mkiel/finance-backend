package middleware

import (
	"github.com/S4mkiel/finance-backend/application/dto"
	"github.com/S4mkiel/finance-backend/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"time"
)

func AuthMiddleware(c *fiber.Ctx) error {
	claims := &utils.Claims{}
	authToken := c.Get(fiber.HeaderAuthorization)

	bearerToken, err := utils.ValidationBearerToken(authToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.BaseDto{
			Success: utils.PBool(false),
			Message: utils.PString("Unauthorized: Invalid token format"),
		})
	}

	token, err := jwt.ParseWithClaims(*bearerToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("bananacomaçucar"), nil
	})
	if err != nil || !token.Valid {
		log.Println("Token parsing error:", err)
		return c.Status(fiber.StatusUnauthorized).JSON(dto.BaseDto{
			Success: utils.PBool(false),
			Message: utils.PString("Unauthorized: Invalid token"),
		})
	}

	c.Locals("userID", claims.UserID)

	return c.Next()
}

func RateLimitMiddleware() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        10,
		Expiration: 30 * time.Second,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(dto.BaseDto{
				Success: utils.PBool(false),
				Message: utils.PString("Too many requests. Try again later."),
			})
		},
	})
}
