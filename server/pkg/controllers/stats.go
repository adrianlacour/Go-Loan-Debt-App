// Package controllers/loan provides an interface between the route and the DB for loan information

package controllers

import (
	"net/http"

	"../models"
	"github.com/gin-gonic/gin"
)

// CalculateLoanStats generates the stats for all of the user's loans
// Route: GET /stats
func CalculateLoanStats(c *gin.Context) {
	var loans []models.Loan
	models.DB.Find(&loans)

	c.JSON(http.StatusOK, gin.H{"data": loans})
}

func calcTotalPrincipal(l *[]models.Loan) float32 {
	return 12.2
}
