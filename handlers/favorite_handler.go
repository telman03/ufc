package handlers

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/telman03/ufc/db"
	"github.com/telman03/ufc/models"
)
// AddFavorite godoc
// @Summary Add a fighter to favorites
// @Description Add a fighter to the authenticated user's favorites list
// @Tags favorites
// @Accept json
// @Produce json
// @Param request body models.FavoriteInput true "Favorite Request"
// @Security BearerAuth
// @Router /favorites [post]
func AddFavorite(c echo.Context) error {
	user := c.Get("user").(jwt.MapClaims)
	userID := uint(user["user_id"].(float64))

	var input models.FavoriteInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	var fighter models.Fighter
	if err := db.DB.First(&fighter, input.FighterID).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Fighter not found"})
	}

	var existingFavorite models.Favorite
	if err := db.DB.Where("user_id = ? AND fighter_id = ?", userID, input.FighterID).First(&existingFavorite).Error; err == nil {
		return c.JSON(http.StatusConflict, echo.Map{"error": "fighter already favorited"})
	}

	favorite := models.Favorite{UserID: userID, FighterID: input.FighterID}
	if err := db.DB.Create(&favorite).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to add favorite"})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "Fighter added to favorites"})

}

// RemoveFavorite godoc
// @Summary Remove a fighter from favorites
// @Description Remove a fighter from the authenticated user's favorites list
// @Tags favorites
// @Produce json
// @Param fighter_id path int true "Fighter ID"
// @Security BearerAuth
// @Router /favorites/{fighter_id} [delete]
func RemoveFavorite(c echo.Context) error {

	user := c.Get("user").(jwt.MapClaims)
	userID := uint(user["user_id"].(float64))


	fighterID := c.Param("fighter_id")


	if err := db.DB.Where("user_id = ? AND fighter_id = ?", userID, fighterID).Delete(&models.Favorite{}).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to remove favorite"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Fighter removed from favorites"})

}

// ListFavorites godoc
// @Summary List favorite fighters
// @Description List all fighters favorited by the authenticated user
// @Tags favorites
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Fighter
// @Router /favorites [get]
func ListFavorites(c echo.Context) error {
	user := c.Get("user").(jwt.MapClaims)
	userID := uint(user["user_id"].(float64))

	var favorites []models.Favorite
	if err := db.DB.Preload("Fighter").Where("user_id = ?", userID).Find(&favorites).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch favorites"})
	}

	var fighters []models.Fighter
	for _, favorite := range favorites {
		fighters = append(fighters, favorite.Fighter)
	}

	return c.JSON(http.StatusOK, fighters)
}