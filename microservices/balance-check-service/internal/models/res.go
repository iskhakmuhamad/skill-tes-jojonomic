package models

type Response struct {
	Error   bool        `json:"error"`
	ReffID  string      `json:"reff_id"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `josn:"data,omitempty"`
}
