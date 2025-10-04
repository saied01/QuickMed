package main

import (
	"quickmed/internal/db"
	//"quickmed/internal/email"
	//"quickmed/internal/reservation"
	"quickmed/internal/user"

	"github.com/gin-gonic/gin"
	//"quickmed/internal/auth"
)

func main() {
	database := db.New()

	db.Truncate(database, "users")

	// database to repository and services
	userRepo := user.NewUserRepository(database)
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)
	// reservationRepo := reservation.NewRepository(database)

	r := gin.Default()

	r.GET("/users/:id", userHandler.GetUser)
	r.POST("/signup", userHandler.SignUp)
	r.DELETE("/users/:id", userHandler.DeleteUser)

	r.Run(":8081")
}
