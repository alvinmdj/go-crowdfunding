package handler

import (
	"go-crowdfunding/helper"
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

func (h *userHandler) RegisterUser(c *gin.Context) {
	// get user input & map user input to RegisterUserInput struct
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	// pass struct as parameter to service
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	// token, err := h.jwtService.GenerateToken()

	formatter := user.FormatUser(newUser, "tokentokentokentokentoken")

	response := helper.APIResponse(
		"Account has been created",
		http.StatusOK,
		"success",
		formatter,
	)

	c.JSON(http.StatusOK, response)
}
