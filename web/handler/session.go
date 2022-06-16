package handler

import (
	"go-crowdfunding/user"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type sessionHandler struct {
	userService user.Service
}

func NewSessionHandler(userService user.Service) *sessionHandler {
	return &sessionHandler{userService}
}

// Handler to show the login page
func (h *sessionHandler) Create(c *gin.Context) {
	c.HTML(http.StatusOK, "session_create.html", nil)
}

// Handler to login the user
func (h *sessionHandler) Store(c *gin.Context) {
	var input user.LoginInput

	// bind & validate the input
	err := c.ShouldBind(&input)
	if err != nil {
		c.Redirect(http.StatusFound, "/")
		return
	}

	// login user
	user, err := h.userService.Login(input)
	if err != nil || user.Role != "admin" {
		c.Redirect(http.StatusFound, "/")
		return
	}

	// set the current user's id & name to session
	session := sessions.Default(c)
	session.Set("userId", user.ID)
	session.Set("userName", user.Name)
	session.Save() // save the session

	c.Redirect(http.StatusFound, "/users")
}

func (h *sessionHandler) Destroy(c *gin.Context) {
	// delete the session
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.Redirect(http.StatusFound, "/")
}
