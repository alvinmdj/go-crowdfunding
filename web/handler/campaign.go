package handler

import (
	"go-crowdfunding/campaign"
	"go-crowdfunding/user"
	"net/http"

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
