package dto

import "github.com/nikzayn/banking/errs"

type TransactionRequest struct {
	AccountId       string  `json:"account_id"`
	CustomerId      string  `json:"-"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}

func (t TransactionRequest) Validate() *errs.AppError {
	if t.TransactionType != "deposit" && t.TransactionType != "withdrawl" {
		return errs.ValidationError("Transaction type can only be deposit or withdrawl")
	}
	if t.Amount < 0 {
		return errs.ValidationError("Amount cannot be less than zero")
	}
	return nil
}
