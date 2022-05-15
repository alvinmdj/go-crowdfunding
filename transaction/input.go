package transaction

// Get campaign transactions input
// Input Campaign ID from URI parameter
type GetCampaignTransactionsInput struct {
	ID int `uri:"id" binding:"required"` // use uri not json
}
