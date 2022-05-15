package transaction

type Service interface {
	GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error)
}

type service struct {
	repository Repository
}

// Transaction service instance
func NewService(repository Repository) *service {
	return &service{repository}
}

// Service to get transactions by campaign ID
func (s *service) GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error) {
	transactions, err := s.repository.FindByCampaignID(input.ID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
