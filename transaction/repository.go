package transaction

import "gorm.io/gorm"

type Repository interface {
	FindByCampaignID(campaignId int) ([]Transaction, error)
	FindByUserID(userId int) ([]Transaction, error)
	FindByCode(transactionCode string) (Transaction, error)
	Create(transaction Transaction) (Transaction, error)
	Update(transaction Transaction) (Transaction, error)
	FindAll() ([]Transaction, error)
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

// Repository to get transaction by transaction code
func (r *repository) FindByCode(transactionCode string) (Transaction, error) {
	var transaction Transaction

	err := r.db.Where("code = ?", transactionCode).Find(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

// Repository to create transaction
func (r *repository) Create(transaction Transaction) (Transaction, error) {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

// Repository to update transaction
func (r *repository) Update(transaction Transaction) (Transaction, error) {
	err := r.db.Save(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

// Repository to get all transactions
func (r *repository) FindAll() ([]Transaction, error) {
	var transactions []Transaction

	err := r.db.Preload("Campaign").Preload("User").Order("id desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
