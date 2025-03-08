package main

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/telman03/ufc/docs"
	"github.com/telman03/ufc/models"
	"github.com/telman03/ufc/scheduler"

	 "github.com/telman03/ufc/scraper"

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
	db.DB.AutoMigrate(&models.User{}, &models.Fighter{}, &models.Favorite{}, &models.Ranking{}, &models.Event{}, &models.Fight{})

	// go scraper.ScrapeAndStoreFighters()
	// go scraper.ScrapeAndStoreRankings()
	 go scraper.ScrapeUpcomingEvents()
	 go scraper.ScrapeFightCards()
	go scheduler.StartScheduler()

	e.POST("/register", handlers.Register)
	e.POST("/login", handlers.Login)

	e.GET("/profile", handlers.ProtectedRoute, middleware.AuthMiddleware)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/fighters", handlers.SearchFighters)
	e.GET("/rankings", handlers.GetRankingsByWeightClass)

	e.POST("/favorites", handlers.AddFavorite, middleware.AuthMiddleware)
	e.DELETE("/favorites/:fighter_id", handlers.RemoveFavorite, middleware.AuthMiddleware)
	e.GET("/favorites", handlers.ListFavorites, middleware.AuthMiddleware)
	
	e.GET("/events/upcoming", handlers.GetUpcomingEvents)
	e.GET("/events/:event_id/fightcard", handlers.GetFightCard)


	e.POST("/admin/events", handlers.CreateEvent, middleware.AdminMiddleware)
	e.PUT("/admin/events/:id", handlers.UpdateEvent, middleware.AdminMiddleware)
	e.DELETE("/admin/events/:id", handlers.DeleteEvent, middleware.AdminMiddleware)
	e.POST("/admin/users/:id/role", handlers.UpdateUserRole, middleware.AdminMiddleware)
	
	e.Logger.Fatal(e.Start(":8080"))
}
