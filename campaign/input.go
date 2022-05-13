package campaign

type GetCampaignDetailsInput struct {
	// uri: api/v1/campaigns/{id}
	ID int `uri:"id" binding:"required"` // use uri not json
}
