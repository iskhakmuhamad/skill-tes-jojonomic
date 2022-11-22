package app

import (
	"encoding/json"
	"net/http"
	"price-check-service/configs"
	"price-check-service/internal/models"
)

func CheckPrice(w http.ResponseWriter, r *http.Request) {
	c := configs.NewConfig()

	m := new(models.Price)

	if err := c.DB.Model(m).Order("created_at DESC").First(m).Error; err != nil {
		sendResponse(w, http.StatusBadRequest, true, err.Error(), nil)
		return
	}

	sendResponse(w, http.StatusOK, false, "", m)
}

func sendResponse(w http.ResponseWriter, code int, isError bool, msg string, data interface{}) {
	rs := &models.Response{
		Error:   isError,
		Message: msg,
		Data:    data,
	}

	resp, err := json.Marshal(rs)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(resp)
}
