package transaction

import "gorm.io/gorm"

type Repository interface {
	FindByCampaignID(campaignId int) ([]Transaction, error)
	FindByUserID(userId int) ([]Transaction, error)
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
		Order("id desc").
		Find(&transactions).
		Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

// Repository to get transactions by user ID
func (r *repository) FindByUserID(userId int) ([]Transaction, error) {
	var transactions []Transaction

	// Campaign.CampaignImages -> CampaignImages is defined inside Campaign struct.
	// Preload Campaign & CampaignImages related to the Campaign,
	// where campaign_images is_primary = 1
	err := r.db.
		Preload("Campaign.CampaignImages", "campaign_images.is_primary = 1").
		Where("user_id = ?", userId).
		Order("id desc").
		Find(&transactions).
		Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}