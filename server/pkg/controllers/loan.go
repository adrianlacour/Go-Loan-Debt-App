// Package controllers/loan provides an interface between the route and the DB for loan information

package controllers

import (
	"net/http"

	"../models"
	"github.com/gin-gonic/gin"
)

// A CreateLoanInput respresents the required input from the user to input a new loan
type CreateLoanInput struct {
	LoanName     string  `json:"loan_name" binding:"required"`
	Debt         float32 `json:"debt" binding:"required"`
	InterestRate float32 `json:"interest_rate" binding:"required"`
}

// CreateLoan inputs a new loan into the database
// Route: POST /loans
func CreateLoan(c *gin.Context) {

	var input CreateLoanInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loan := models.Loan{LoanName: input.LoanName, Debt: input.Debt, InterestRate: input.InterestRate}
	models.DB.Create(&loan)

	c.JSON(http.StatusOK, gin.H{"data": loan})
}

// FindLoans connects to the database and retrieves all loans inside of it
// Route: GET /loans
func FindLoans(c *gin.Context) {
	var loans []models.Loan
	models.DB.Find(&loans)

	c.JSON(http.StatusOK, gin.H{"data": loans})
}
