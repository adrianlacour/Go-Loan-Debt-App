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

	if len(loans) == 0 {
		c.JSON(http.StatusOK, gin.H{"data": nil})
	}

	var stats models.Stats
	calcLoanLifeStates(&loans, &stats)

	c.JSON(http.StatusOK, gin.H{"data": loans})
}

func calcLoanLifeStates(loans *[]models.Loan, stats *models.Stats) {
	// TODO: Move all stats to using the model
	stats.TotalPrincipal = 0.0
	for i, l := range *loans {
		stats.TotalPrincipal += l.Debt
		stats.AvgInterestRate += (stats.AvgInterestRate + l.InterestRate) / float32(i)
	}

	calcPayoffDate(loans, stats)
}

func calcPayoffDate(loans *[]models.Loan, stats *models.Stats) {
}
