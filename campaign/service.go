package campaign

import (
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(userId int) ([]Campaign, error)
	GetCampaignByID(input GetCampaignDetailsInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

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
