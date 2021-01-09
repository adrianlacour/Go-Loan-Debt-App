// Package models/stats provides base models for loan stats

package models

// A Stats respresents the stats of all loans for a user
type Stats struct {
	TotalPrincipal  float32 `json:"total_principal""`
	PaidOffDate     string  `json:"paid_off_date"`
	TotalInterest   float32 `json:"total_interest"`
	AvgInterestRate float32 `json:"avg_interest_rate"`
}
