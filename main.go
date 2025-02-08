package main

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/telman03/ufc/docs"
	"github.com/telman03/ufc/models"
	// "github.com/telman03/ufc/scraper"

	// "gorm.io/gorm"

	// "github.com/telman03/ufc/scraper"

	"github.com/telman03/ufc/db"
	"github.com/telman03/ufc/handlers"
	"github.com/telman03/ufc/middleware"
)

// @title UFC API
// @version 1.0
// @description This is an API for UFC Fight tracking
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	e := echo.New()

	e.Use(middleware.CORS())

	db.ConnectDB()
	db.DB.AutoMigrate(&models.User{}, &models.Fighter{}, &models.Favorite{}, &models.Ranking{})

	// go scraper.ScrapeAndStoreFighters()
	// go scraper.ScrapeAndStoreRankings()
	e.POST("/register", handlers.Register)
	e.POST("/login", handlers.Login)

	e.GET("/profile", handlers.ProtectedRoute, middleware.AuthMiddleware)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/fighters", handlers.SearchFighters)

	e.Logger.Fatal(e.Start(":8080"))
}
