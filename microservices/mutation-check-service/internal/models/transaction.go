package models

type Transaction struct {
	ReffID string `json:"reff_id" gorm:"type:varchar(15);primaryKey"`
	TransactionData
	CreatedAt int64 `json:"created_at" gorm:"type:integer;autoCreateTime"`
	UpdatedAt int64 `json:"updated_at" gorm:"type:integer;autoUpdateTime"`
}

type TransactionData struct {
	Type         string  `json:"type" gorm:"type:varchar(20)"`
	Gram         float64 `json:"gram" gorm:"type:decimal(12,3)"`
	Balance      float64 `json:"balance" gorm:"type:decimal(12,3)"`
	PriceBuyback float64 `json:"price_buyback" gorm:"type:decimal(12,3)"`
	PriceTopup   float64 `json:"price_topup" gorm:"type:decimal(12,3);"`
	AccountNo    string  `json:"account_no" gorm:"type:varchar(20)"`
}
