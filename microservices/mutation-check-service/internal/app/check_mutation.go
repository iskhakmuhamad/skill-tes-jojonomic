package app

import (
	"encoding/json"
	"net/http"

	"mutation-check-service/configs"
	"mutation-check-service/internal/models"
)

func CheckMutation(w http.ResponseWriter, r *http.Request) {
	c := configs.NewConfig()

	p := new(models.CheckMutationRequest)

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(p); err != nil {
		sendResponse(w, http.StatusBadRequest, true, ``, err.Error(), nil)
		return
	}

	m := make([]*models.Transaction, 0)

	if err := c.DB.Model(m).Where("account_no = ? AND created_at >= ? AND created_at <= ?", p.AccountNo, p.StartDate, p.EndDate).Find(&m).Error; err != nil {
		sendResponse(w, http.StatusBadRequest, true, ``, err.Error(), nil)
		return
	}

	if len(m) == 0 {
		sendResponse(w, http.StatusNotFound, true, ``, "there nothing data transaction", nil)
		return
	}

	sendResponse(w, http.StatusOK, false, ``, ``, m)
}

func sendResponse(w http.ResponseWriter, code int, isError bool, id string, msg string, data interface{}) {
	rs := &models.Response{
		Error:   isError,
		ReffID:  id,
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
