package handler

import (
	"go-crowdfunding/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

// Handler to show list of users page
func (h *userHandler) Index(c *gin.Context) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "user_index.html", gin.H{"users": users})
}

// Handler to show create user page
func (h *userHandler) Create(c *gin.Context) {
	c.HTML(http.StatusOK, "user_create.html", nil)
}

// Handler to store the newly created user
func (h *userHandler) Store(c *gin.Context) {
	var input user.FormRegisterUserInput

	// Get input from form & validate it
	err := c.ShouldBind(&input)
	if err != nil {
		input.Error = err
		c.HTML(http.StatusOK, "user_create.html", input)
		return
	}

	// convert FormRegisterUserInput to RegisterUserInput
	registerInput := user.RegisterUserInput{}
	registerInput.Name = input.Name
	registerInput.Occupation = input.Occupation
	registerInput.Email = input.Email
	registerInput.Password = input.Password

	// call service to register a new user with RegisterUserInput
	_, err = h.userService.RegisterUser(registerInput)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/users")
}
