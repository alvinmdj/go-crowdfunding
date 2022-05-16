package transaction

import (
	"errors"
	"go-crowdfunding/campaign"
)

type Service interface {
	GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error)
	GetTransactionsByUserID(userId int) ([]Transaction, error)
}

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
}

// Transaction service instance
func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository, campaignRepository}
}

// Service to get transactions by campaign ID
func (s *service) GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error) {
	// get campaign
	campaign, err := s.campaignRepository.FindByID(input.ID)
	if err != nil {
		return []Transaction{}, err
	}

	// check if current user is owner of campaign
	if campaign.UserID != input.User.ID {
		return []Transaction{}, errors.New("you are not authorized to access this campaign")
	}

	// get transactions
	transactions, err := s.repository.FindByCampaignID(input.ID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

// Service to get transactions by user ID
func (s *service) GetTransactionsByUserID(userId int) ([]Transaction, error) {
	// get transactions
	transactions, err := s.repository.FindByUserID(userId)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
