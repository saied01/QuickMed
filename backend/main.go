package main

import (
	"quickmed/internal/db"
	//"quickmed/internal/email"
	//"quickmed/internal/reservation"
	"html/template"
	"net/http"
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

	r.Run(":8081")
}
