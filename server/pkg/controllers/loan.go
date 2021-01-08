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

// A UpdateLoanInput respresents the input from the user to update an existing loan
type UpdateLoanInput struct {
	ID           uint    `json:"-"`
	LoanName     string  `json:"loan_name"`
	Debt         float32 `json:"debt"`
	InterestRate float32 `json:"interest_rate"`
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

// FindLoan connects to the database and retrieves a single loan, by id
// Route: GET /loans/:id
func FindLoan(c *gin.Context) { // Get model if exist
	var loan models.Loan

	if err := models.DB.Where("id = ?", c.Param("id")).First(&loan).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": loan})
}

// UpdateLoan updates a loan by id. All fields are optional
// Route: PUT /loans/:id
func UpdateLoan(c *gin.Context) {
	var loan models.Loan
	if err := models.DB.Where("id = ?", c.Param("id")).First(&loan).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateLoanInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&loan).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": loan})
}

// DeleteLoan deletes the loan entry from the db
// Route: DELETE /loan/:id
func DeleteLoan(c *gin.Context) {
	var loan models.Loan
	if err := models.DB.Where("id = ?", c.Param("id")).First(&loan).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&loan)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
