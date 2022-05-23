package payment

import (
	"go-crowdfunding/helper"
	"go-crowdfunding/user"
	"os"

	"github.com/veritrans/go-midtrans"
)

type Service interface {
	GetPaymentURL(transaction Transaction, user user.User) (string, error)
}

type service struct{}

func NewService() *service {
	return &service{}
}

// Midtrans Snap Gateway (https://github.com/veritrans/go-midtrans)
func (s *service) GetPaymentURL(transaction Transaction, user user.User) (string, error) {
	helper.LoadEnv()
	MIDTRANS_SERVER_KEY := os.Getenv("MIDTRANS_SERVER_KEY")
	MIDTRANS_CLIENT_KEY := os.Getenv("MIDTRANS_CLIENT_KEY")

	midclient := midtrans.NewClient()
	midclient.ServerKey = MIDTRANS_SERVER_KEY
	midclient.ClientKey = MIDTRANS_CLIENT_KEY
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  transaction.Code,
			GrossAmt: int64(transaction.Amount),
		},
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Name,
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}

	return snapTokenResp.RedirectURL, nil
}
