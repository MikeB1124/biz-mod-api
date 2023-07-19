package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MikeB1124/biz-mod-api/controllers"
	"github.com/apex/gateway"
	"github.com/gin-gonic/gin"
)

func inLambda() bool {
	if lambdaTaskRoot := os.Getenv("LAMBDA_TASK_ROOT"); lambdaTaskRoot != "" {
		return true
	}
	return false
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	transactions := router.Group("/transactions")
	{
		transactions.GET("/", controllers.GetAllTransactions)
		transactions.GET("/:id", controllers.GetTransactionByID)
		transactions.GET("/date", controllers.GetDatePeriodOfTransactions)
		transactions.POST("/create", controllers.CreateTransaction)
		transactions.DELETE("/delete", controllers.DeleteTransaction)
	}
	return router
}

func main() {
	if inLambda() {
		fmt.Println("running aws lambda in aws")
		log.Fatal(gateway.ListenAndServe(":8080", setupRouter()))
	} else {
		fmt.Println("running aws lambda in local")
		log.Fatal(http.ListenAndServe(":8080", setupRouter()))
	}
}
