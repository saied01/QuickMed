package user

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *UserService
}

func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{service: service}
}

// GET api/users/:id
func (h *UserHandler) GetUserJSON(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
	}

	user, err := h.service.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GET /users/:id
func (h *UserHandler) GetUserPage(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid id")
		return
	}

	user, err := h.service.GetUserByID(uint(id))
	if err != nil {
		c.String(http.StatusNotFound, "user not found")
		return
	}

	tmpl := template.Must(template.ParseFiles(
		"templates/layout.html",
		"templates/user.html",
	))
	c.Status(http.StatusOK)
	tmpl.ExecuteTemplate(c.Writer, "layout", gin.H{
		"Title": "Perfil de Usuario",
		"User":  user,
	})
}

// GET /signup
func (h *UserHandler) SignUpPage(c *gin.Context) {
	tmpl := template.Must(template.ParseFiles(
		"templates/layout.html",
		"templates/signup.html",
	))
	c.Status(http.StatusOK)
	tmpl.ExecuteTemplate(c.Writer, "layout", gin.H{"Title": "Registro"})
}

// POST /signup
func (h *UserHandler) SignUp(c *gin.Context) {

	name := c.PostForm("name")
	email := c.PostForm("email")
	// agestr := c.PostForm("age")
	password := c.PostForm("password")

	if name == "" || email == "" || password == "" {
		c.String(http.StatusBadRequest, "Todos los campos son obligatorios")
		return
	}

	// age, _ := strconv.ParseUint(agestr, 10, 8)

	u, err := h.service.RegisterUser(email, name, password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tmpl := template.Must(template.ParseFiles(
		"templates/user.html",
	))
	tmpl.ExecuteTemplate(c.Writer, "content", gin.H{"User": u})
}

// DELETE /user/:id
func (h *UserHandler) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
	}

	err = h.service.DeleteUser(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "user deleted"})
}
