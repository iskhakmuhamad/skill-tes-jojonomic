package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"buyback-input-service/configs"
	"buyback-input-service/internal/models"

	"github.com/segmentio/kafka-go"
	"github.com/teris-io/shortid"
	"gorm.io/gorm"
)

func Buyback(w http.ResponseWriter, r *http.Request) {
	c := configs.NewConfig()
	p := new(models.BuybackRequest)

	id, _ := shortid.Generate()

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		sendResponse(w, http.StatusBadRequest, true, id, err.Error())
		return
	}

	account, err := getAccount(c.DB, p.AccountNo)
	if err != nil {
		sendResponse(w, http.StatusBadRequest, true, id, err.Error())
		return
	}

	price, err := getCurrentPrice(c.DB)
	if err != nil {
		sendResponse(w, http.StatusBadRequest, true, id, err.Error())
		return
	}

	if p.Gram > account.Balance {
		sendResponse(w, http.StatusBadRequest, true, id, "your gold balance not enaught")
		return
	}

	m := &models.BuybackData{
		ReffID:       id,
		Gram:         p.Gram,
		Price:        p.Price,
		AccountNo:    p.AccountNo,
		Balance:      account.Balance,
		PriceTopup:   price.PriceTopup,
		PriceBuyback: price.PriceBuyback,
	}

	defer r.Body.Close()

	pByte, err := json.Marshal(&m)
	if err != nil {
		sendResponse(w, http.StatusBadRequest, true, id, err.Error())
		return
	}

	c.Kafka.SetWriteDeadline(time.Now().Add(10 * time.Second))

	msg := kafka.Message{
		Key:   []byte(fmt.Sprintf("address-%s", r.RemoteAddr)),
		Value: pByte,
	}

	_, err = c.Kafka.WriteMessages(msg)
	if err != nil {
		log.Printf("error write message kafka with error message : %s\n", err.Error())
		sendResponse(w, http.StatusBadRequest, true, id, "Kafka not ready")
		return
	}

	sendResponse(w, http.StatusCreated, false, id, ``)
}

func getAccount(db *gorm.DB, accountNo string) (*models.Account, error) {
	var account *models.Account
	if err := db.Model(account).Where("account_no = ?", accountNo).First(&account).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("account not found")
		}
		return nil, err
	}

	return account, nil
}

func getCurrentPrice(db *gorm.DB) (*models.Price, error) {
	var h *models.Price
	if err := db.Model(h).Order("created_at DESC").First(&h).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("price not found")
		}
		return nil, err
	}

	return h, nil
}

func sendResponse(w http.ResponseWriter, code int, isError bool, id string, msg string) {
	rs := &models.Response{
		Error:   isError,
		ReffID:  id,
		Message: msg,
	}

	resp, err := json.Marshal(rs)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(resp)
}
