package main

import (
	"fmt"
	"log"
	"net/http"

	"price-input-service/configs"
	"price-input-service/internal/app"
)

func main() {
	c := configs.NewConfig()
	router := c.Router

	router.HandleFunc("/api/input-harga", app.PriceInput).Methods(http.MethodPost)

	log.Printf("api running in port %d", c.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", c.Port), router)
}
