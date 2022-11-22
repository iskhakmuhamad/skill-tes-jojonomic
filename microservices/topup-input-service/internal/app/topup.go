package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"topup-input-service/configs"
	"topup-input-service/internal/models"

	"github.com/segmentio/kafka-go"
	"github.com/shopspring/decimal"
	"github.com/teris-io/shortid"
	"gorm.io/gorm"
)

func Topup(w http.ResponseWriter, r *http.Request) {
	c := configs.NewConfig()
	var data models.TopupData

	id, _ := shortid.Generate()

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		sendResponse(w, http.StatusBadRequest, true, id, err.Error())
		return
	}

	m := &models.Topup{
		ReffID:    id,
		TopupData: data,
	}

	defer r.Body.Close()

	pByte, err := json.Marshal(&m)
	if err != nil {
		sendResponse(w, http.StatusBadRequest, true, id, err.Error())
		return
	}

	if err := validateTopup(c.DB, &m.TopupData); err != nil {
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

func validateTopup(db *gorm.DB, data *models.TopupData) error {
	m := new(models.Price)

	if err := db.Model(m).Order("created_at DESC").First(&m).Error; err != nil {
		return err
	}

	price, err := strconv.ParseFloat(data.Price, 64)
	if err != nil {
		return err
	}

	if price != m.PriceData.PriceTopup {
		return errors.New("price doesn't match with current price topup")
	}

	gram, err := strconv.ParseFloat(data.Gram, 64)
	if err != nil {
		return err
	}

	// cekmultiplied :=
	if !decimal.NewFromFloat(gram*1000).BigFloat().IsInt() || gram <= 0 {
		return errors.New("minimum gram top-up is multiply by 0.001")
	}

	return nil
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
