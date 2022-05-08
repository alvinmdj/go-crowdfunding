package handler

import (
	"fmt"
	"go-crowdfunding/auth"
	"go-crowdfunding/helper"
	"go-crowdfunding/user"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// get user input
	var input user.RegisterUserInput

	// validate user input
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors} // gin.H is a map (key - value)

		response := helper.APIResponse(
			"Failed to create user", http.StatusUnprocessableEntity, "error", errorMessage,
		)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// pass struct as parameter to service
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse(
			"Failed to create user", http.StatusBadRequest, "error", nil,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := helper.APIResponse(
			"Failed to create user", http.StatusBadRequest, "error", nil,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, token)

	response := helper.APIResponse(
		"Account has been created", http.StatusOK, "success", formatter,
	)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) LoginUser(c *gin.Context) {
	// get user input (email & password) & map user input to LoginUserInput struct
	var input user.LoginInput

	// validate user input
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors} // gin.H is a map (key - value)

		response := helper.APIResponse(
			"Login failed", http.StatusUnprocessableEntity, "error", errorMessage,
		)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// check if user exists & check if password is correct
	loggedInUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse(
			"Login failed", http.StatusUnprocessableEntity, "error", errorMessage,
		)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.authService.GenerateToken(loggedInUser.ID)
	if err != nil {
		response := helper.APIResponse(
			"Login failed", http.StatusBadRequest, "error", nil,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(loggedInUser, token)

	response := helper.APIResponse(
		"Login successful", http.StatusOK, "success", formatter,
	)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	// get user input (email)
	var input user.CheckEmailInput

	// validate user input
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse(
			"Failed to check email availability", http.StatusUnprocessableEntity, "error", errorMessage,
		)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// check if input email is available
	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"error": "Internal server error"}
		response := helper.APIResponse(
			"Failed to check email availability", http.StatusBadRequest, "error", errorMessage,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_available": isEmailAvailable}

	metaMessage := "This email is already taken"
	if isEmailAvailable {
		metaMessage = "Email is available"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	// get input from user (image) -> Form Data, not JSON
	file, err := c.FormFile("avatar") // avatar is the name of the input
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse(
			"Failed to upload avatar", http.StatusBadRequest, "error", data,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// get current user from context (from auth middleware)
	currentUser := c.MustGet("currentUser").(user.User)
	userId := currentUser.ID

	path := fmt.Sprintf("images/%d-%d-%s", userId, time.Now().UnixMilli(), file.Filename)

	// save image in folder 'images/'
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse(
			"Failed to upload avatar", http.StatusBadRequest, "error", data,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// update user avatar in database
	_, err = h.userService.SaveAvatar(userId, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse(
			"Failed to upload avatar", http.StatusBadRequest, "error", data,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse(
		"Avatar has been uploaded", http.StatusOK, "success", data,
	)
	c.JSON(http.StatusOK, response)
}
