package models

type Price struct {
	ReffID string `json:"reff_id"`
	PriceData
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

type PriceData struct {
	PriceTopup   float64 `json:"price_topup"`
	PriceBuyback float64 `json:"price_buyback"`
	AdminID      string  `json:"admin_id"`
}
