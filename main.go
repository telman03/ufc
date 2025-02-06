package main

import (

	"github.com/labstack/echo/v4"
	"github.com/telman03/ufc/db"
	"github.com/telman03/ufc/handlers"
	"github.com/telman03/ufc/models"

)

func main() {
	e := echo.New()

	db.ConnectDB()
	db.DB.AutoMigrate(&models.User{}, &models.Fighter{}, &models.Favorite{})


	e.POST("/register", handlers.Register)
	e.POST("/login", handlers.Login)
	
	e.GET("/profile", handlers.ProtectedRoute, AuthMiddleware)

	e.Logger.Fatal(e.Start(":8080"))
}

