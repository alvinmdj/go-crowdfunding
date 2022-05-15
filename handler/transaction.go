package handler

import (
	"go-crowdfunding/helper"
	"go-crowdfunding/transaction"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	transactionService transaction.Service
}

// Transaction handler instance
func NewTransactionHandler(transactionService transaction.Service) *transactionHandler {
	return &transactionHandler{transactionService}
}

// Handler to get transactions by campaign ID
func (h *transactionHandler) GetCampaignTransactions(c *gin.Context) {
	var input transaction.GetCampaignTransactionsInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse(
			"Failed to get campaign transactions", http.StatusBadRequest, "error", nil,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	transactions, err := h.transactionService.GetTransactionsByCampaignID(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse(
			"Failed to get campaign transactions", http.StatusBadRequest, "error", errorMessage,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// formatter := transaction.FormatCampaignTransactions(transactions)

	response := helper.APIResponse(
		"Campaign transactions", http.StatusOK, "success", transactions,
	)
	c.JSON(http.StatusOK, response)
}
