package transaction

import "gorm.io/gorm"

type Repository interface {
	FindByCampaignID(campaignId int) ([]Transaction, error)
}

type repository struct {
	db *gorm.DB
}

// Transaction repository instance
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// Repository to get transactions by campaign ID
func (r *repository) FindByCampaignID(campaignId int) ([]Transaction, error) {
	var transactions []Transaction

	err := r.db.
		Preload("User").
		Where("campaign_id = ?", campaignId).
		Order("created_at desc").
		Find(&transactions).
		Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
