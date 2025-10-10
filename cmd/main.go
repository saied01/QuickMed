package main

import (
	"quickmed/internal/db"
	"quickmed/internal/reservation"
	//"quickmed/internal/email"
	"html/template"
	"net/http"
	"quickmed/internal/user"

	"github.com/gin-gonic/gin"
	//"quickmed/internal/auth"
)

func main() {
	database := db.New()

	db.Truncate(database, "users")
	// db.Truncate(database, "reservations")

	// database to repository and services
	userRepo := user.NewUserRepository(database)
	userService := user.NewUserService(userRepo)
	reservationRepo := reservation.NewReservationRepository(database)
	reservationService := reservation.NewReservationService(reservationRepo)
	resHandler := reservation.NewReservationHandler(reservationService)
	userHandler := user.NewUserHandler(userService, reservationService)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		tmpl := template.Must(template.ParseFiles(
			"templates/layout.html",
			"templates/index.html",
		))
		c.Status(http.StatusOK)
		tmpl.ExecuteTemplate(c.Writer, "layout", gin.H{
			"Title": "Inicio",
		})
	})

	// API JSON
	r.GET("/api/users/:id", userHandler.GetUserJSON)
	r.DELETE("/users/:id", userHandler.DeleteUser)

	// Web HTML (htmx)
	r.GET("/users/:id", userHandler.GetUserPage)
	r.POST("/signup", userHandler.SignUp)
	r.GET("/signup", userHandler.SignUpPage)
	r.POST("/reservations", resHandler.CreateReservation)

	r.Run(":8081")
}
