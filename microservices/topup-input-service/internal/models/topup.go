package models

type Topup struct {
	ReffID string `json:"reff_id"`
	TopupData
}

type TopupData struct {
	Gram      string `json:"gram"`
	Price     string `json:"price"`
	AccountNo string `json:"account_no"`
}
