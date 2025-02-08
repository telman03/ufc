package handlers

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/telman03/ufc/db"
	"github.com/telman03/ufc/models"
	"net/http"
)


// SearchFighters godoc
// @Summary Search and filter fighters
// @Description Search fighters based on name, stance, weight, wins, losses, and paginate results
// @Tags Fighters
// @Accept json
// @Produce json
// @Param name query string false "Fighter's Name"
// @Param stance query string false "Fighter's Stance (e.g., Southpaw, Orthodox)"
// @Param weight query string false "Fighter's Weight"
// @Param wins query int false "Minimum Wins"
// @Param losses query int false "Maximum Losses"
// @Param limit query int false "Limit number of results (default is 10)"
// @Param offset query int false "Offset for pagination (default is 0)"
// @Success 200 {array} models.Fighter "List of fighters"
// @Failure 400 {object} map[string]interface{} "Invalid query parameters"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /fighters [get]
func SearchFighters(c echo.Context) error {
	var fighters []models.Fighter
	query := db.DB

	if name:= c.QueryParam("name"); name !=""{
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}

	  if stance := c.QueryParam("stance"); stance != "" {
        query = query.Where("stance = ?", stance)
    }

    if weight := c.QueryParam("weight"); weight != "" {
        query = query.Where("weight = ?", weight)
    }

    if wins := c.QueryParam("wins"); wins != "" {
        query = query.Where("wins >= ?", wins)
    }

    if losses := c.QueryParam("losses"); losses != "" {
        query = query.Where("losses <= ?", losses)
    }

    limit := 10 // Default limit
    if c.QueryParam("limit") != "" {
        fmt.Sscanf(c.QueryParam("limit"), "%d", &limit)
    }
    offset := 0
    if c.QueryParam("offset") != "" {
        fmt.Sscanf(c.QueryParam("offset"), "%d", &offset)
    }

    query = query.Limit(limit).Offset(offset)

    if err := query.Find(&fighters).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch fighters"})
    }

    return c.JSON(http.StatusOK, fighters)
}