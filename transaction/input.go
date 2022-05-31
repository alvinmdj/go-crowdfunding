package transaction

import "go-crowdfunding/user"

// Get campaign transactions input
// Input Campaign ID from URI parameter
type GetCampaignTransactionsInput struct {
	ID   int `uri:"id" binding:"required"` // use uri not json
	User user.User
}

// Create transaction input
type CreateTransactionInput struct {
	Amount     int `json:"amount" binding:"required"`
	CampaignID int `json:"campaign_id" binding:"required"`
	User       user.User
}

// Transaction notification input
// https://docs.midtrans.com/en/after-payment/http-notification
type TransactionNotificationInput struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}
