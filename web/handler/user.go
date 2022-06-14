package handler

import (
	"fmt"
	"go-crowdfunding/user"
	"net/http"
	"strconv"
	"time"

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

// Handler to show edit user page
func (h *userHandler) Edit(c *gin.Context) {
	// get id from url
	idFromParam := c.Param("id")

	// convert idFromParam (which is a string) to int
	userId, err := strconv.Atoi(idFromParam)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	// call service to get user by id
	registeredUser, err := h.userService.GetUserById(userId)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	// map registeredUser to FormUpdateUserInput
	formUpdateUserInput := user.FormUpdateUserInput{}
	formUpdateUserInput.ID = registeredUser.ID
	formUpdateUserInput.Name = registeredUser.Name
	formUpdateUserInput.Occupation = registeredUser.Occupation
	formUpdateUserInput.Email = registeredUser.Email

	c.HTML(http.StatusOK, "user_edit.html", formUpdateUserInput)
}

// Handler to update the user
func (h *userHandler) Update(c *gin.Context) {
	idFromParam := c.Param("id")

	// convert idFromParam (which is a string) to int
	userId, err := strconv.Atoi(idFromParam)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	// Get input from form & validate it
	var input user.FormUpdateUserInput
	err = c.ShouldBind(&input)
	if err != nil {
		input.ID = userId
		input.Error = err
		c.HTML(http.StatusOK, "user_edit.html", input)
		return
	}

	// bind user id to input
	input.ID = userId

	// call service to update user
	_, err = h.userService.UpdateUser(input)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/users")
}

// Handler to show upload avatar page
func (h *userHandler) NewAvatar(c *gin.Context) {
	idFromParam := c.Param("id")

	userId, err := strconv.Atoi(idFromParam)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "user_upload.html", gin.H{"ID": userId})
}

// Handler to upload user avatar
func (h *userHandler) StoreAvatar(c *gin.Context) {
	idFromParam := c.Param("id")

	userId, err := strconv.Atoi(idFromParam)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	// Get file input from form & validate it
	file, err := c.FormFile("avatar")
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	// generate unique number from current time in milli
	uniqueNumber := time.Now().UnixMilli()

	// save image in folder 'public/images/avatars/'
	rootPath := fmt.Sprintf("public/images/avatars/%d-%d-%s", userId, uniqueNumber, file.Filename)
	err = c.SaveUploadedFile(file, rootPath)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	// update user avatar in database (path: avatars/filename.extension)
	relativePath := fmt.Sprintf("avatars/%d-%d-%s", userId, uniqueNumber, file.Filename)
	_, err = h.userService.SaveAvatar(userId, relativePath)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/users")
}
