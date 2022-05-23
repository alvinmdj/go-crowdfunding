package transaction

import (
	"errors"
	"fmt"
	"time"

	"go-crowdfunding/campaign"
	"go-crowdfunding/payment"
)

type Service interface {
	GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error)
	GetTransactionsByUserID(userId int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
}

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
	paymentService     payment.Service
}

// Transaction service instance
func NewService(repository Repository, campaignRepository campaign.Repository, paymentService payment.Service) *service {
	return &service{repository, campaignRepository, paymentService}
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

	// map new transaction to payment.Transaction struct
	paymentTransaction := payment.Transaction{
		Code:   newTransaction.Code,
		Amount: newTransaction.Amount,
	}

	// call payment service to get payment URL
	paymentUrl, err := s.paymentService.GetPaymentURL(paymentTransaction, input.User)
	if err != nil {
		return newTransaction, err
	}

	// update transaction with payment URL
	newTransaction.PaymentURL = paymentUrl
	newTransaction, err = s.repository.Update(newTransaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}
