package campaign

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Campaign, error)
	FindWithLimit(limit int) ([]Campaign, error)
	FindByUserID(userId int) ([]Campaign, error)
	FindByID(id int) (Campaign, error)
	FindBySlug(slug string) (Campaign, error)
	Create(campaign Campaign) (Campaign, error)
	Update(campaign Campaign) (Campaign, error)
	CreateImage(campaignImage CampaignImage) (CampaignImage, error)
	MarkAllImagesAsNotPrimary(campaignId int) (bool, error)
}

type repository struct {
	db *gorm.DB
}

// Campaign repository instance
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// Repository to get all campaigns
func (r *repository) FindAll() ([]Campaign, error) {
	var campaigns []Campaign

	// Preload "CampaignImages" -> from Campaign struct entity, where campaign_images table's is_primary = 1
	err := r.db.
		Preload("CampaignImages", "campaign_images.is_primary = 1").
		Order("id desc").
		Find(&campaigns).
		Error

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

// Repository to get campaigns with limit
func (r *repository) FindWithLimit(limit int) ([]Campaign, error) {
	var campaigns []Campaign

	// Preload "CampaignImages" -> from Campaign struct entity, where campaign_images table's is_primary = 1
	err := r.db.
		Preload("CampaignImages", "campaign_images.is_primary = 1").
		Order("id desc").
		Limit(limit).
		Find(&campaigns).
		Error

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

// Repository to get campaigns by user ID
func (r *repository) FindByUserID(userId int) ([]Campaign, error) {
	var campaigns []Campaign

	err := r.db.
		Preload("CampaignImages", "campaign_images.is_primary = 1").
		Where("user_id = ?", userId).
		Order("id desc").
		Find(&campaigns).
		Error

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

// Repository to get campaign details
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

// Repository to check if slug is unique
func (r *repository) FindBySlug(slug string) (Campaign, error) {
	var campaign Campaign

	err := r.db.Where("slug = ?", slug).Find(&campaign).Error

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

// Repository to create a new campaign
func (r *repository) Create(campaign Campaign) (Campaign, error) {
	err := r.db.Create(&campaign).Error

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

// Repository to update a campaign
func (r *repository) Update(campaign Campaign) (Campaign, error) {
	err := r.db.Save(&campaign).Error

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

// Repository to create a new campaign image
func (r *repository) CreateImage(campaignImage CampaignImage) (CampaignImage, error) {
	err := r.db.Create(&campaignImage).Error

	if err != nil {
		return campaignImage, err
	}

	return campaignImage, nil
}

// Repository to mark all images as not primary
func (r *repository) MarkAllImagesAsNotPrimary(campaignId int) (bool, error) {
	// UPDATE campaign_images SET is_primary = 0 WHERE campaign_id = ?
	err := r.db.Model(&CampaignImage{}).Where("campaign_id = ?", campaignId).Update("is_primary", false).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

// ! Repository to delete a campaign
// ! Currently not in use
// func (r *repository) Delete(id int) error {
// 	err := r.db.Delete(&Campaign{}, "id = ?", id).Error

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
