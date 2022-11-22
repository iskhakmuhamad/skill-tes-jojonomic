package models

type BuybackData struct {
	ReffID       string  `json:"reff_id"`
	Gram         float64 `json:"gram"`
	Price        float64 `json:"price"`
	AccountNo    string  `json:"account_no"`
	PriceTopup   float64 `json:"price_topup"`
	PriceBuyback float64 `json:"price_buyback"`
	Balance      float64 `json:"balance"`
}
