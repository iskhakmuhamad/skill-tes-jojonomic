package models

type CheckMutationRequest struct {
	AccountNo string `json:"account_no"`
	StartDate int32  `json:"start_date"`
	EndDate   int32  `json:"end_date"`
}
