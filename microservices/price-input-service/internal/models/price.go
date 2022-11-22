package models

type Price struct {
	ReffID string `json:"reff_id"`
	PriceData
}

type PriceData struct {
	PriceTopup   float64 `json:"price_topup"`
	PriceBuyback float64 `json:"price_buyback"`
	AdminID      string  `json:"admin_id"`
}
