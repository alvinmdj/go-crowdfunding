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

// Create campaign image input
// Input from form data
type CreateCampaignImageInput struct {
	CampaignID int       `form:"campaign_id" binding:"required"`
	IsPrimary  bool      `form:"is_primary"`
	User       user.User // from jwt
}

// FormCreateCampaignInput is the input required for creating a campaign
// Here, we use form, not json
type FormCreateCampaignInput struct {
	Name             string `form:"name" binding:"required"`
	ShortDescription string `form:"short_description" binding:"required"`
	Description      string `form:"description" binding:"required"`
	GoalAmount       int    `form:"goal_amount" binding:"required"`
	Perks            string `form:"perks" binding:"required"`
	UserID           int    `form:"user_id" binding:"required"`
	Users            []user.User
	Error            error
}

// FormUpdateCampaignInput is the input required for updating a campaign
type FormUpdateCampaignInput struct {
	ID               int
	Name             string `form:"name" binding:"required"`
	ShortDescription string `form:"short_description" binding:"required"`
	Description      string `form:"description" binding:"required"`
	GoalAmount       int    `form:"goal_amount" binding:"required"`
	Perks            string `form:"perks" binding:"required"`
	Error            error
	User             user.User
}
