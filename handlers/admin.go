package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/telman03/ufc/db"
	"github.com/telman03/ufc/models"
)

// CreateEvent godoc
// @Summary Create a new event
// @Description Admins can create a new event
// @Tags admin
// @Accept  json
// @Produce  json
// @Param request body models.Event true "Event Details"
// @Security BearerAuth
// @Router /admin/events [post]
func CreateEvent(c echo.Context) error {
	var event models.Event
	if err := c.Bind(&event); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	if err := db.DB.Create(&event).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create event"})
	}

	return c.JSON(http.StatusCreated, event)
}

// UpdateEvent godoc
// @Summary Update an existing event
// @Description Admins can update an event by ID
// @Tags admin
// @Accept  json
// @Produce  json
// @Param id path int true "Event ID"
// @Param request body models.Event true "Updated Event Details"
// @Security BearerAuth
// @Router /admin/events/{id} [put]
func UpdateEvent(c echo.Context) error {
	var event models.Event
	id := c.Param("id")

	if err := db.DB.First(&event, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Event not found"})
	}

	if err := c.Bind(&event); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	if err := db.DB.Save(&event).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to update event"})
	}

	return c.JSON(http.StatusOK, event)
}


// DeleteEvent godoc
// @Summary Delete an event
// @Description Admins can delete an event by ID
// @Tags admin
// @Produce  json
// @Param id path int true "Event ID"
// @Security BearerAuth
// @Router /admin/events/{id} [delete]
func DeleteEvent(c echo.Context) error {
	var event models.Event
	id := c.Param("id")

	if err := db.DB.Delete(&event, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to delete event"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Event deleted successfully"})
}

// UpdateUserRole godoc
// @Summary Update user role
// @Description Admins can change a user's role (e.g., user → admin)
// @Tags admin
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param request body object{role=string} true "New Role (admin/user)"
// @Security BearerAuth
// @Router /admin/users/{id}/role [post]
func UpdateUserRole(c echo.Context) error {
	var user models.User
	id := c.Param("id")

	if err := db.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "User not found"})
	}

	var request struct {
		Role string `json:"role"`
	}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	user.Role = request.Role
	if err := db.DB.Save(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to update role"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "User role updated successfully"})
}
