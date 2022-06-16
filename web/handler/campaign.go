package handler

import (
	"fmt"
	"go-crowdfunding/campaign"
	"go-crowdfunding/user"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	campaignService campaign.Service
	userService     user.Service
}

func NewCampaignHandler(campaignService campaign.Service, userService user.Service) *campaignHandler {
	return &campaignHandler{campaignService, userService}
}

// Handler to show list of campaigns page
func (h *campaignHandler) Index(c *gin.Context) {
	// get all campaigns
	campaigns, err := h.campaignService.GetCampaigns(0, 0)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "campaign_index.html", gin.H{"campaigns": campaigns})
}

// Handler to show create campaign page
func (h *campaignHandler) Create(c *gin.Context) {
	// get all users
	users, err := h.userService.GetAllUsers()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	input := campaign.FormCreateCampaignInput{}
	input.Users = users

	c.HTML(http.StatusOK, "campaign_create.html", input)
}

// Handler to create campaign
func (h *campaignHandler) Store(c *gin.Context) {
	var input campaign.FormCreateCampaignInput

	// bind & validate input
	err := c.ShouldBind(&input)
	if err != nil {
		// get users again for the select user dropdown
		users, e := h.userService.GetAllUsers()
		if e != nil {
			c.HTML(http.StatusInternalServerError, "error.html", nil)
			return
		}

		// set users & error
		input.Users = users
		input.Error = err

		c.HTML(http.StatusOK, "campaign_create.html", input)
		return
	}

	// get user by id to map to CreateCampaignInput
	user, err := h.userService.GetUserById(input.UserID)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	// map form input to CreateCampaignInput
	createCampaignInput := campaign.CreateCampaignInput{}
	createCampaignInput.Name = input.Name
	createCampaignInput.ShortDescription = input.ShortDescription
	createCampaignInput.Description = input.Description
	createCampaignInput.GoalAmount = input.GoalAmount
	createCampaignInput.Perks = input.Perks
	createCampaignInput.User = user

	// call service to create campaign
	_, err = h.campaignService.CreateCampaign(createCampaignInput)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/campaigns")
}

// Handler to show create campaign image form
func (h *campaignHandler) NewImage(c *gin.Context) {
	// get campaign id from uri
	idFromParam := c.Param("id")

	id, err := strconv.Atoi(idFromParam)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "campaign_upload.html", gin.H{"ID": id})
}

// Handler to create campaign image
func (h *campaignHandler) StoreImage(c *gin.Context) {
	// get file from form data
	file, err := c.FormFile("file")
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	// get campaign id from uri
	idFromParam := c.Param("id")
	campaignId, err := strconv.Atoi(idFromParam)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	// get user id from current campaign
	existingCampaign, err := h.campaignService.GetCampaignByID(campaign.GetCampaignDetailsInput{ID: campaignId})
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	userId := existingCampaign.UserID

	// generate unique number from current time in milli
	uniqueNumber := time.Now().UnixMilli()

	// save campaign image in folder 'public/images/campaign-images/'
	rootPath := fmt.Sprintf("public/images/campaign-images/%d-%d-%s", userId, uniqueNumber, file.Filename)
	err = c.SaveUploadedFile(file, rootPath)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	// get user by id to map to CreateCampaignInput
	userCampaign, err := h.userService.GetUserById(userId)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	// map required input to CreateCampaignImageInput
	createCampaignImageInput := campaign.CreateCampaignImageInput{}
	createCampaignImageInput.CampaignID = campaignId
	createCampaignImageInput.IsPrimary = true
	createCampaignImageInput.User = userCampaign

	// save campaign image in database (path: campaign-images/filename.extension)
	relativePath := fmt.Sprintf("campaign-images/%d-%d-%s", userId, uniqueNumber, file.Filename)
	_, err = h.campaignService.CreateCampaignImage(createCampaignImageInput, relativePath)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/campaigns")
}

// Handler to show edit campaign form
func (h *campaignHandler) Edit(c *gin.Context) {
	// get campaign id from uri
	idFromParam := c.Param("id")

	// convert id to int
	id, err := strconv.Atoi(idFromParam)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	// get campaign by id
	existingCampaign, err := h.campaignService.GetCampaignByID(campaign.GetCampaignDetailsInput{ID: id})
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	// map campaign data to FormUpdateCampaignInput
	input := campaign.FormUpdateCampaignInput{}
	input.ID = id
	input.Name = existingCampaign.Name
	input.ShortDescription = existingCampaign.ShortDescription
	input.Description = existingCampaign.Description
	input.GoalAmount = existingCampaign.GoalAmount
	input.Perks = existingCampaign.Perks

	c.HTML(http.StatusOK, "campaign_edit.html", input)
}

// Handler to update campaign
func (h *campaignHandler) Update(c *gin.Context) {
	idFromParam := c.Param("id")

	id, err := strconv.Atoi(idFromParam)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	// bind & validate input
	var input campaign.FormUpdateCampaignInput
	err = c.ShouldBind(&input)
	if err != nil {
		input.Error = err
		input.ID = id
		c.HTML(http.StatusOK, "campaign_edit.html", input)
		return
	}

	// get campaign by id
	existingCampaign, err := h.campaignService.GetCampaignByID(campaign.GetCampaignDetailsInput{ID: id})
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	userId := existingCampaign.UserID

	// get user by id to map to CreateCampaignInput
	userCampaign, err := h.userService.GetUserById(userId)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	// map form input to UpdateCampaignInput
	updateInput := campaign.CreateCampaignInput{}
	updateInput.Name = input.Name
	updateInput.ShortDescription = input.ShortDescription
	updateInput.Description = input.Description
	updateInput.GoalAmount = input.GoalAmount
	updateInput.Perks = input.Perks
	updateInput.User = userCampaign

	_, err = h.campaignService.UpdateCampaign(campaign.GetCampaignDetailsInput{ID: id}, updateInput)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/campaigns")
}

// Handler to show campaign details
func (h *campaignHandler) Show(c *gin.Context) {
	// get campaign id from uri
	idFromParam := c.Param("id")

	// convert id to int
	id, err := strconv.Atoi(idFromParam)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	// get campaign by id
	existingCampaign, err := h.campaignService.GetCampaignByID(campaign.GetCampaignDetailsInput{ID: id})
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "campaign_show.html", existingCampaign)
}
