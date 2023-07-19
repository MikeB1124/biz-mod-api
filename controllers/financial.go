package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllTransactions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Function to fetch all sales and expenses."})
}

func GetTransactionByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Function to get a sepcific transaction."})
}

func GetDatePeriodOfTransactions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Function to get transactions from a date period."})
}

func CreateTransaction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Function to create sale or expense."})
}

func DeleteTransaction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Function to delete a sepcific transaction."})
}
