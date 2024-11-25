package routes

import (
	"github.com/VieiraGabrielAlexandre/fms_vieira/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Rotas para transações
	transactionRoutes := r.Group("/transactions")
	{
		transactionRoutes.GET("/credit", controllers.GetCredits)
		transactionRoutes.POST("/credit", controllers.CreateCredit)
		transactionRoutes.GET("/debit", controllers.GetDebits)
		transactionRoutes.POST("/debit", controllers.CreateDebit)
	}
}
