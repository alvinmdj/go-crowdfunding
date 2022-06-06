package handler

import (
	"fmt"
	"go-crowdfunding/campaign"
	"go-crowdfunding/helper"
	"go-crowdfunding/user"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	campaignService campaign.Service
}

// Campaign handler instance
func NewCampaignHandler(campaignService campaign.Service) *campaignHandler {
	return &campaignHandler{campaignService}
}

// Handler to get all campaigns
func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	campaigns, err := h.campaignService.GetCampaigns(userId, limit)
	if err != nil {
		response := helper.APIResponse(
			"Error trying to get campaigns data", http.StatusBadRequest, "error", nil,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := campaign.FormatCampaigns(campaigns)

	response := helper.APIResponse(
		"List of campaigns", http.StatusOK, "success", formatter,
	)
	c.JSON(http.StatusOK, response)
}

// Handler to get campaign details
func (h *campaignHandler) GetCampaignDetails(c *gin.Context) {
	var input campaign.GetCampaignDetailsInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse(
			"Failed to get campaign details", http.StatusBadRequest, "error", nil,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignDetails, err := h.campaignService.GetCampaignByID(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse(
			"Failed to get campaign details", http.StatusBadRequest, "error", errorMessage,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := campaign.FormatCampaignDetails(campaignDetails)

	response := helper.APIResponse(
		"Campaign details", http.StatusOK, "success", formatter,
	)
	c.JSON(http.StatusOK, response)
}

// Handler to create a new campaign
func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var input campaign.CreateCampaignInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse(
			"Failed to create campaign", http.StatusUnprocessableEntity, "error", errorMessage,
		)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// get current user from context
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	// call service to create campaign
	newCampaign, err := h.campaignService.CreateCampaign(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse(
			"Failed to create campaign", http.StatusBadRequest, "error", errorMessage,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := campaign.FormatCampaign(newCampaign)
	response := helper.APIResponse(
		"Campaign created", http.StatusOK, "success", formatter,
	)
	c.JSON(http.StatusOK, response)
}

// Handler to update a campaign
func (h *campaignHandler) UpdateCampaign(c *gin.Context) {
	// get campaign id from uri
	var inputId campaign.GetCampaignDetailsInput
	err := c.ShouldBindUri(&inputId)
	if err != nil {
		response := helper.APIResponse(
			"Failed to update campaign", http.StatusBadRequest, "error", nil,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// get campaign data from body
	var inputData campaign.CreateCampaignInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse(
			"Failed to update campaign", http.StatusUnprocessableEntity, "error", errorMessage,
		)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// get current user from context
	currentUser := c.MustGet("currentUser").(user.User)
	inputData.User = currentUser

	// call service to update campaign
	updatedCampaign, err := h.campaignService.UpdateCampaign(inputId, inputData)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse(
			"Failed to update campaign", http.StatusBadRequest, "error", errorMessage,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := campaign.FormatCampaign(updatedCampaign)
	response := helper.APIResponse(
		"Campaign updated", http.StatusOK, "success", formatter,
	)
	c.JSON(http.StatusOK, response)
}

// Handler to upload campaign image
func (h *campaignHandler) UploadCampaignImage(c *gin.Context) {
	var input campaign.CreateCampaignImageInput

	err := c.ShouldBind(&input) // use 'c.ShouldBind' to bind form data
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(
			"Failed to upload campaign image", http.StatusUnprocessableEntity, "error", errorMessage,
		)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// get current user from context (from auth middleware)
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser // add current user to input
	userId := currentUser.ID // for image name

	// get image input from Form Data, not JSON
	file, err := c.FormFile("file") // file is the name of the input
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse(
			"Failed to upload campaign image", http.StatusBadRequest, "error", data,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// save campaign image in folder 'public/images/campaign-images/'
	rootPath := fmt.Sprintf("public/images/campaign-images/%d-%d-%s", userId, time.Now().UnixMilli(), file.Filename)
	err = c.SaveUploadedFile(file, rootPath)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse(
			"Failed to upload campaign image", http.StatusBadRequest, "error", data,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// save campaign image in database (path: campaign-images/filename.extension)
	relativePath := fmt.Sprintf("campaign-images/%d-%d-%s", userId, time.Now().UnixMilli(), file.Filename)
	_, err = h.campaignService.CreateCampaignImage(input, relativePath)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse(
			"Failed to upload campaign image", http.StatusBadRequest, "error", data,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse(
		"Campaign image has been uploaded", http.StatusOK, "success", data,
	)
	c.JSON(http.StatusOK, response)
}
