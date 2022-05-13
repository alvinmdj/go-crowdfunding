package campaign

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Campaign, error)
	FindByUserID(userId int) ([]Campaign, error)
	FindByID(id int) (Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Campaign, error) {
	var campaigns []Campaign

	// Preload "CampaignImages" -> from Campaign struct entity, where campaign_images table's is_primary = 1
	err := r.db.
		Preload("CampaignImages", "campaign_images.is_primary = 1").
		Find(&campaigns).
		Error

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindByUserID(userId int) ([]Campaign, error) {
	var campaigns []Campaign

	err := r.db.
		Preload("CampaignImages", "campaign_images.is_primary = 1").
		Where("user_id = ?", userId).
		Find(&campaigns).
		Error

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindByID(id int) (Campaign, error) {
	var campaign Campaign

	err := r.db.
		Preload("User").
		Preload("CampaignImages").
		Where("id = ?", id).
		Find(&campaign).
		Error

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}
