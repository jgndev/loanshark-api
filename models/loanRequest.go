package models

type LoanRequest struct {
	Amount float64 `json:"amount"`
	Rate   float64 `json:"rate"`
	Term   int     `json:"term"`
}
