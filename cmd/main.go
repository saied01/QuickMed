package main

import (
	"log"

	"quickmed/internal/db"
	//"quickmed/internal/email"
	//"quickmed/internal/reservation"
	"quickmed/internal/user"
	//"quickmed/internal/auth"
)

func main() {
	database := db.New()

	if err := database.AutoMigrate(&user.User{}); err != nil {
		log.Fatal(err)
	}

	// database to repository and services
	// userRepo := user.NewRepository(database)
	// userService := user.NewService(userRepo)
	// reservationRepo := reservation.NewRepository(database)
}
