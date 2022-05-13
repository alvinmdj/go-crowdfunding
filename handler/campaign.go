package handler

import (
	"go-crowdfunding/campaign"
	"go-crowdfunding/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	campaignService campaign.Service
}

func NewCampaignHandler(campaignService campaign.Service) *campaignHandler {
	return &campaignHandler{campaignService}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.campaignService.GetCampaigns(userId)
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

func (h *campaignHandler) GetCampaignDetails(c *gin.Context) {
	// api/v1/campaings/{id}
	// handler : mapping id from url to input struct -> service, call formatter
	// service : input struct -> get id from url -> call repo
	// repository : get campaign by id

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
