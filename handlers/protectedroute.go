package handlers

import (
	"net/http"


	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func ProtectedRoute(c echo.Context) error {
	// Get claims from context
	claims := c.Get("user").(jwt.MapClaims) // Access as MapClaims, not *jwt.Token

	// Now you can access claims data like:
	userID := claims["user_id"]
	username := claims["username"]

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Protected route accessed successfully",
		"username": username,
		"user_id": userID,
	})
}
