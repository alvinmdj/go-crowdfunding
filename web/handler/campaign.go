package handler

import (
	"go-crowdfunding/campaign"
	"net/http"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	campaignService campaign.Service
}

func NewCampaignHandler(campaignService campaign.Service) *campaignHandler {
	return &campaignHandler{campaignService}
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
