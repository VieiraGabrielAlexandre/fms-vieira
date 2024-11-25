package controllers

import (
	"net/http"

	"github.com/VieiraGabrielAlexandre/fms_vieira/services"

	"github.com/VieiraGabrielAlexandre/fms_vieira/models"

	"github.com/gin-gonic/gin"
)

// GET /transactions/credit
func GetCredits(c *gin.Context) {
	// Capture os dois valores retornados
	credits, err := services.GetCredits()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"credits": credits})
}

// POST /transactions/credit
func CreateCredit(c *gin.Context) {
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction.Type = "credit"
	if err := services.CreateTransaction(transaction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Credit created successfully"})
}

// GET /transactions/debit
func GetDebits(c *gin.Context) {
	// Capture os dois valores retornados
	debits, err := services.GetDebits()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"debits": debits})
}

// POST /transactions/debit
func CreateDebit(c *gin.Context) {
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction.Type = "debit"
	if err := services.CreateTransaction(transaction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Debit created successfully"})
}
