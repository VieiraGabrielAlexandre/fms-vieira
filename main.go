package main

import (
	"github.com/VieiraGabrielAlexandre/fms_vieira/routes"

	"github.com/VieiraGabrielAlexandre/fms_vieira/database"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Carregar variáveis de ambiente
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Conectar ao banco de dados
	database.ConnectDatabase()

	// Configurar rotas
	r := gin.Default()
	routes.SetupRoutes(r)

	// Iniciar o servidor
	r.Run(":8080")
}
