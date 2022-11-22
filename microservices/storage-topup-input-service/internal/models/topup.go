package models

type Topup struct {
	Reff_ID string `json:"reff_id" gorm:"type:varchar(15);primaryKey"`
	TopupData
	CreatedAt int64 `json:"created_at" gorm:"type:integer;autoCreateTime"`
	UpdatedAt int64 `json:"updated_at" gorm:"type:integer;autoUpdateTime"`
}

type TopupData struct {
	Gram      string `json:"gram" gorm:"type:varchar(100)"`
	Price     string `json:"price" gorm:"type:varchar(100)"`
	AccountNo string `json:"account_no" gorm:"type:varchar(20)"`
}
