package models

type LoanResponse struct {
	Amount         string    `json:"loanAmount"`
	Rate           string    `json:"loanRate"`
	MonthlyPayment string    `json:"monthlyPayment"`
	TotalInterest  string    `json:"totalInterest"`
	TotalCost      string    `json:"totalCost"`
	Payments       []Payment `json:"payments"`
}
