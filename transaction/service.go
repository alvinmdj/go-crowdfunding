package transaction

import (
	"errors"
	"fmt"
	"time"

	"go-crowdfunding/campaign"
)

type Service interface {
	GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error)
	GetTransactionsByUserID(userId int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
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

// Service to create transaction
func (s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	// map input to Transaction struct
	transaction := Transaction{}
	transaction.Amount = input.Amount
	transaction.CampaignID = input.CampaignID
	transaction.UserID = input.User.ID
	transaction.Status = "pending"
	transaction.Code = fmt.Sprintf("TRX-%d%d%d", input.CampaignID, input.User.ID, time.Now().UnixMilli())

	// call repository to create transaction
	newTransaction, err := s.repository.Create(transaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}
