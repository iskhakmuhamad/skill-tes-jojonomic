package models

type Price struct {
	ReffID string `json:"reff_id" gorm:"type:varchar(15);primaryKey"`
	PriceData
	CreatedAt int64 `json:"created_at" gorm:"type:integer;autoCreateTime"`
	UpdatedAt int64 `json:"updated_at" gorm:"type:integer;autoUpdateTime"`
}

type PriceData struct {
	PriceTopup   float64 `json:"price_topup" gorm:"type:decimal(12,3);"`
	PriceBuyback float64 `json:"price_buyback" gorm:"type:decimal(12,3)"`
	AdminID      string  `json:"admin_id" gorm:"type:varchar(15);"`
}
