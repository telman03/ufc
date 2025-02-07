package handlers

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

// ProtectedRoute godoc
// @Summary Get user profile
// @Description Retrieve the authenticated user's profile using JWT token
// @Tags auth
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "User profile data"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Router /profile [get]
func ProtectedRoute(c echo.Context) error {
	claims := c.Get("user").(jwt.MapClaims)

	userID := claims["user_id"]
	username := claims["username"]

	return c.JSON(http.StatusOK, echo.Map{
		"message":  "Protected route accessed successfully",
		"username": username,
		"user_id":  userID,
	})
}
