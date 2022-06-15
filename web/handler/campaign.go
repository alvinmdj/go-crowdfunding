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
