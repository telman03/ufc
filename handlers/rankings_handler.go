package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/telman03/ufc/db"
	"github.com/telman03/ufc/models"
)



type FighterRankingResponse struct {
	Division string `json:"division"`
	Rank     int    `json:"rank"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Wins     int    `json:"wins"`
	Losses   int    `json:"losses"`
	Draws    int    `json:"draws"`
}
// GetRankingsByWeightClass retrieves fighter rankings filtered by weight class
// @Summary Get rankings by weight class
// @Description Retrieve the rankings of fighters for a specific weight class
// @Tags Rankings
// @Accept json
// @Produce json
// @Param weightclass query string true "Weight class of fighters (e.g., Lightweight, Welterweight)"
// @Success 200 {array} Fighter "List of ranked fighters"
// @Failure 400 {object} map[string]interface{} "Bad Request - Missing weight class parameter"
// @Failure 500 {object} map[string]interface{} "Internal Server Error - Database query issue"
// @Router /rankings [get]
func GetRankingsByWeightClass(c echo.Context) error {
	weightClass := c.QueryParam("weightclass")

	if weightClass == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Weight class is required"})
	}

	var rankings []models.Ranking
	result := db.DB.Preload("Fighter").
		Where("division = ?", weightClass).
		Order("rank ASC").
		Find(&rankings)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Database error"})
	}

	var response []FighterRankingResponse
	for _, ranking := range rankings {
		response = append(response, FighterRankingResponse{
			Division: ranking.Division,
			Rank:     ranking.Rank,
			Name:     ranking.Fighter.Name,
			Nickname: ranking.Fighter.Nickname,
			Wins:     ranking.Fighter.Wins,
			Losses:   ranking.Fighter.Losses,
			Draws:    ranking.Fighter.Draws,
		})
	}

	return c.JSON(http.StatusOK, response)
}
