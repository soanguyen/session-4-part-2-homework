package auth

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// JWTMiddleware creates a middleware that validates JWT tokens
func JWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Get the Authorization header
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "missing authorization header")
			}

			// Check if the header has the Bearer prefix
			if !strings.HasPrefix(authHeader, "Bearer ") {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid token format")
			}

			// Extract the token
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			// Verify the token
			username, err := VerifyToken(tokenString)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
			}

			// Set the username in the context for later use
			c.Set("username", username)

			// Continue with the next middleware or handler
			return next(c)
		}
	}
}
