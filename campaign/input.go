package campaign

import "go-crowdfunding/user"

// Get campaign details input
// Input Campaign ID from URI parameter
type GetCampaignDetailsInput struct {
	// uri: api/v1/campaigns/{id}
	ID int `uri:"id" binding:"required"` // use uri not json
}

// Create campaign input
type CreateCampaignInput struct {
	Name             string    `json:"name" binding:"required"`
	ShortDescription string    `json:"short_description" binding:"required"`
	Description      string    `json:"description" binding:"required"`
	GoalAmount       int       `json:"goal_amount" binding:"required"`
	Perks            string    `json:"perks" binding:"required"`
	User             user.User // from jwt
}
