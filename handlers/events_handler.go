package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/telman03/ufc/db"
	"github.com/telman03/ufc/models"
)



// GetUpcomingEvents retrieves all upcoming UFC events
// @Summary Get upcoming UFC events
// @Description Fetches a list of upcoming UFC events
// @Tags events
// @Produce json
// @Success 200 {array} models.Event
// @Router /events/upcoming [get]
func GetUpcomingEvents(c echo.Context) error {
	var events []models.Event
	if err:= db.DB.Order("id ASC").Find(&events).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch events"})
	}
	return c.JSON(http.StatusOK, events)
}

// GetFightCard retrieves the fight card for a given event
// @Summary Get fight card for an event
// @Description Fetches all fights scheduled for a specific UFC event
// @Tags events
// @Produce json
// @Param event_id path int true "Event ID"
// @Success 200 {array} models.Fight
// @Router /events/{event_id}/fightcard [get]
func GetFightCard(c echo.Context) error {
	eventID := c.Param("event_id")

	var fights []models.Fight

	if err := db.DB.Where("event_id = ?", eventID).Find(&fights).Error; err!=nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch events"})
	}
	return c.JSON(http.StatusOK, fights)
}