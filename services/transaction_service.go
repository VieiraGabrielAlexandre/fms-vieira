package services

import (
	"github.com/VieiraGabrielAlexandre/fms_vieira/database"
	"github.com/VieiraGabrielAlexandre/fms_vieira/models"
)

func GetCredits() ([]models.Transaction, error) {
	var credits []models.Transaction
	err := database.DB.Where("type = ?", "credit").Find(&credits).Error
	return credits, err
}

func GetDebits() ([]models.Transaction, error) {
	var debits []models.Transaction
	err := database.DB.Where("type = ?", "debit").Find(&debits).Error
	return debits, err
}

func CreateTransaction(transaction models.Transaction) error {
	return database.DB.Create(&transaction).Error
}
