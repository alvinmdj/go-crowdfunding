package campaign

import "go-crowdfunding/user"

type GetCampaignDetailsInput struct {
	// uri: api/v1/campaigns/{id}
	ID int `uri:"id" binding:"required"` // use uri not json
}

type CreateCampaignInput struct {
	Name             string    `json:"name"`
	ShortDescription string    `json:"short_description"`
	Description      string    `json:"description"`
	GoalAmount       int       `json:"goal_amount"`
	Perks            string    `json:"perks"`
	User             user.User // from jwt
}
