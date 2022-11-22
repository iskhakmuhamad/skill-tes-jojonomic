package models

type BuybackRequest struct {
	Gram      float64 `json:"gram"`
	Price     float64 `json:"price"`
	AccountNo string  `json:"account_no"`
}
