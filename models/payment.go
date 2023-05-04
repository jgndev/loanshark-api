package models

type Payment struct {
	Month     int    `json:"month"`
	Payment   string `json:"payment"`
	Principal string `json:"principal"`
	Interest  string `json:"interest"`
	Balance   string `json:"balance"`
}
