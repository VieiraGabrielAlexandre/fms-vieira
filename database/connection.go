package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	// Configuração de conexão
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),     // Usuário
		os.Getenv("DB_PASSWORD"), // Senha
		os.Getenv("DB_HOST"),     // Host (e.g., localhost)
		os.Getenv("DB_PORT"),     // Porta (e.g., 3306)
		os.Getenv("DB_NAME"),     // Nome do banco
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Migração do modelo
	err = DB.AutoMigrate(&Transaction{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}

// Modelo usado no banco
type Transaction struct {
	gorm.Model
	Amount float64 `json:"amount"`
	Type   string  `json:"type"`
}
