package middlewares //Middleware validates JWT before passing control to the next handler.

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/riz254/Go-Ticketing-System.git/models"
	"gorm.io/gorm"
)

func AuthProtected(db *gorm.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		authHeader := ctx.Get("Authorization")

		if authHeader == "" {
			log.Warnf("empty authorization header")

			return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
				"status":  "fail",
				"message": "unauthorized",
			})
		}

		//token parts
		tokenParts := strings.Split(authHeader, " ") //Splits the authHeader string into parts wherever a space (" ") is found.

		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			log.Warnf("invalid token parts")

			return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
				"status":  "fail",
				"message": "unauthorized",
			})
		}

		tokenStr := tokenParts[1]
		secret := []byte(os.Getenv("JWT_SECRET"))

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if token.Method.Alg() != jwt.GetSigningMethod("HS256").Alg() {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return secret, nil
		})

		//check token validity
		if err != nil || !token.Valid {
			log.Warnf("invalid token")

			return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
				"status":  "fail",
				"message": "unauthorized",
			})
		}

		//if everything is okay get user with token
		userId := token.Claims.(jwt.MapClaims)["id"]

		if err := db.Model(&models.User{}).Where("id = ?", userId).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warnf("user not in db")

			return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
				"status":  "fail",
				"message": "unauthorized",
			})
		}

		ctx.Locals("userId", userId)

		return ctx.Next()
	}
}
