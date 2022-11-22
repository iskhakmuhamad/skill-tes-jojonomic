package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"price-input-service/configs"
	"price-input-service/internal/models"

	"github.com/segmentio/kafka-go"
	"github.com/teris-io/shortid"
)

func PriceInput(w http.ResponseWriter, r *http.Request) {
	c := configs.NewConfig()
	var data models.PriceData

	id, _ := shortid.Generate()

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		sendResponse(w, http.StatusBadRequest, true, id, err.Error())
		return
	}

	m := &models.Price{
		ReffID:    id,
		PriceData: data,
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

	sendResponse(w, http.StatusCreated, false, id, "")
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
