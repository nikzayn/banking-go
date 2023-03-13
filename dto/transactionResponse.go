package dto

import "time"

type TransactionResponse struct {
	TransactionId   string    `json:"transaction_id"`
	AccountId       string    `json:"account_id"`
	Balance         float64   `json:"new_balance"`
	TransactionType string    `json:"transaction_type"`
	TransactionDate time.Time `json:"transaction_date"`
}
