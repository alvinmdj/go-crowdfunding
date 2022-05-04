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
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors} // gin.H is a map (key - value)

		response := helper.APIResponse(
			"Failed to create user",
			http.StatusUnprocessableEntity,
			"error",
			errorMessage,
		)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// pass struct as parameter to service
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse(
			"Failed to create user",
			http.StatusBadRequest,
			"error",
			nil,
		)
		c.JSON(http.StatusBadRequest, response)
		return
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
