package models

type Account struct {
	ReffID string `json:"reff_id" gorm:"type:varchar(15);primaryKey"`
	AccountData
	CreatedAt int64 `json:"created_at" gorm:"type:integer;autoCreateTime"`
	UpdatedAt int64 `json:"updated_at" gorm:"type:integer;autoUpdateTime"`
}

type AccountData struct {
	AccountNo string  `json:"account_no" gorm:"type:varchar(20);unique"`
	Balance   float64 `json:"balance" gorm:"type:decimal(12,3)"`
}
