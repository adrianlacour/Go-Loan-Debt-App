// Package models/loan provides base models for the loan debt information

package models

// A Loan respresents a single loan of the user
type Loan struct {
	ID           uint    `json:"id" gorm:"primaryKey"`
	LoanName     string  `json:"loan_name"`
	Debt         float32 `json:"debt"`
	InterestRate float32 `json:"interest_rate"`
	MonthlyMin   float32 `json:"monthly_min_payment"`
}
