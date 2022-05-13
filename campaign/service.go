package campaign

import "fmt"

type Service interface {
	GetCampaigns(userId int) ([]Campaign, error)
	GetCampaignByID(input GetCampaignDetailsInput) (Campaign, error)
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
