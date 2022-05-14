package campaign

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(userId int) ([]Campaign, error)
	GetCampaignByID(input GetCampaignDetailsInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	UpdateCampaign(inputId GetCampaignDetailsInput, inputData CreateCampaignInput) (Campaign, error)
	CreateCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error)
}

type service struct {
	repository Repository
}

// Campaign service instance
func NewService(repository Repository) *service {
	return &service{repository}
}

// Service to get all campaigns
// If User ID is set, only return campaigns that belong to that user
func (s *service) GetCampaigns(userId int) ([]Campaign, error) {
	if userId == 0 { // no userId, return all campaigns
		campaigns, err := s.repository.FindAll()
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}

	campaigns, err := s.repository.FindByUserID(userId)
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

// Service to get campaign details
func (s *service) GetCampaignByID(input GetCampaignDetailsInput) (Campaign, error) {
	campaign, err := s.repository.FindByID(input.ID)

	if err != nil {
		return campaign, err
	}

	if campaign.ID == 0 {
		return campaign, fmt.Errorf("Campaign with ID %d cannot be found", input.ID)
	}

	return campaign, nil
}

// Service to create a new campaign
func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	// map input to Campaign struct
	campaign := Campaign{}
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.GoalAmount = input.GoalAmount
	campaign.Perks = input.Perks
	campaign.UserID = input.User.ID

	// create slug
	slugCandidate := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	campaign.Slug = slug.Make(slugCandidate)

	// check if campaign with same slug already exists
	existingCampaign, err := s.repository.FindBySlug(campaign.Slug)

	if err != nil {
		return existingCampaign, err
	}

	if existingCampaign.ID != 0 {
		return existingCampaign, fmt.Errorf("Campaign with slug %s already exists", campaign.Slug)
	}

	newCampaign, err := s.repository.Create(campaign)
	if err != nil {
		return newCampaign, err
	}

	return newCampaign, nil
}

// Service to update a campaign
func (s *service) UpdateCampaign(inputId GetCampaignDetailsInput, inputData CreateCampaignInput) (Campaign, error) {
	// get campaign by ID
	campaign, err := s.repository.FindByID(inputId.ID)
	if err != nil {
		return campaign, err
	}

	// check if current user is the owner of the campaign
	if campaign.UserID != inputData.User.ID {
		return campaign, errors.New("unauthorized to update this campaign")
	}

	// map input to Campaign struct
	campaign.Name = inputData.Name
	campaign.ShortDescription = inputData.ShortDescription
	campaign.Description = inputData.Description
	campaign.GoalAmount = inputData.GoalAmount
	campaign.Perks = inputData.Perks

	updatedCampaign, err := s.repository.Update(campaign)
	if err != nil {
		return updatedCampaign, err
	}

	return updatedCampaign, nil
}

// Service to create a new campaign image
func (s *service) CreateCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error) {
	// check if campaign image is primary
	isPrimary := 0
	if input.IsPrimary {
		isPrimary = 1
		_, err := s.repository.MarkAllImagesAsNotPrimary(input.CampaignID)
		if err != nil {
			return CampaignImage{}, err
		}
	}

	// map input to CampaignImage struct
	campaignImage := CampaignImage{}
	campaignImage.CampaignID = input.CampaignID
	campaignImage.IsPrimary = isPrimary
	campaignImage.FileName = fileLocation

	newCampaignImage, err := s.repository.CreateImage(campaignImage)
	if err != nil {
		return newCampaignImage, err
	}

	return newCampaignImage, nil
}
