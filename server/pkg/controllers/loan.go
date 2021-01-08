// Package controllers/loan provides an interface between the route and the DB for loan information

package controllers

import (
	"net/http"

	"../models"
	"github.com/gin-gonic/gin"
)

// FindLoans connects to the database and retrieves all loans inside of it
// Route: GET /loans
func FindLoans(c *gin.Context) {
	var loans []models.Loan
	models.DB.Find(&loans)

	c.JSON(http.StatusOK, gin.H{"data": loans})
}
