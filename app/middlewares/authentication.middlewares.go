package middlewares

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	common_response "github.com/phucpham-infinity/go-nextpress/app/common/response"
	"github.com/phucpham-infinity/go-nextpress/app/context"
)

func AuthenticationMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return common_response.NewAppError(c).IsUnauthorized(errors.New("missing authorization header")).SendJSON()
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return common_response.NewAppError(c).IsUnauthorized(errors.New("invalid authorization header format")).SendJSON()
		}
		tokenString := parts[1]
		jwtConfig := context.AppContext().GetConfig().JWT

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(jwtConfig.Secret), nil
		})
		if err != nil {
			return common_response.NewAppError(c).IsUnauthorized(errors.New("invalid or expired token")).SendJSON()
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return common_response.NewAppError(c).IsUnauthorized(errors.New("invalid token claims")).SendJSON()
		}

		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				return common_response.NewAppError(c).IsUnauthorized(errors.New("token has expired")).SendJSON()
			}
		}
		userID, ok := claims["id"].(string)
		if !ok || userID == "" {
			return common_response.NewAppError(c).IsUnauthorized(errors.New("missing or invalid user id in token")).SendJSON()
		}
		c.Locals("userId", userID)

		userEmail, ok := claims["email"].(string)
		if !ok || userEmail == "" {
			return common_response.NewAppError(c).IsUnauthorized(errors.New("missing or invalid user email in token")).SendJSON()
		}
		c.Locals("userEmail", userEmail)
		return c.Next()
	}
}
