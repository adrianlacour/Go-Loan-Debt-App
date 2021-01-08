// Package models/loan provides base models for the loan debt information

package models

// A Loan respresents a single loan of the user
type Loan struct {
	ID           uint    `json:"id" gorm:"primary_key"`
	LoanName     string  `json:"loan_name"`
	Debt         string  `json:"debt"`
	InterestRate float32 `json:"interest_rate"`
}
