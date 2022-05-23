package handler

import (
	"go-crowdfunding/helper"
	"go-crowdfunding/transaction"
	"go-crowdfunding/user"
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

	// get campaign ID from URI parameter and map to input struct
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse(
			"Failed to get campaign transactions", http.StatusBadRequest, "error", nil,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// get current user from context
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	// call service to get transactions
	transactions, err := h.transactionService.GetTransactionsByCampaignID(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse(
			"Failed to get campaign transactions", http.StatusBadRequest, "error", errorMessage,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := transaction.FormatCampaignTransactions(transactions)

	response := helper.APIResponse(
		"Campaign transactions", http.StatusOK, "success", formatter,
	)
	c.JSON(http.StatusOK, response)
}

// Handler to get transactions by user ID
func (h *transactionHandler) GetUserTransactions(c *gin.Context) {
	// get current user from context
	currentUser := c.MustGet("currentUser").(user.User)
	userId := currentUser.ID

	// call service to get transactions
	transactions, err := h.transactionService.GetTransactionsByUserID(userId)
	if err != nil {
		response := helper.APIResponse(
			"Failed to get user transactions", http.StatusBadRequest, "error", nil,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := transaction.FormatUserTransactions(transactions)

	response := helper.APIResponse(
		"User transactions", http.StatusOK, "success", formatter,
	)
	c.JSON(http.StatusOK, response)
}

// Handler to create transaction
func (h *transactionHandler) CreateTransaction(c *gin.Context) {
	var input transaction.CreateTransactionInput

	// get json data and map to input struct
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(
			"Failed to create transaction", http.StatusUnprocessableEntity, "error", errorMessage,
		)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// get current user from context
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	// call service to create transaction
	newTransaction, err := h.transactionService.CreateTransaction(input)
	if err != nil {
		response := helper.APIResponse(
			"Failed to create transaction", http.StatusBadRequest, "error", nil,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := transaction.FormatTransaction(newTransaction)
	response := helper.APIResponse(
		"Transaction created", http.StatusOK, "success", formatter,
	)
	c.JSON(http.StatusOK, response)
}
