package main

import (
	"fmt"
	"log"
	"net/http"

	"price-check-service/configs"
	"price-check-service/internal/app"
)

func main() {
	c := configs.NewConfig()
	router := c.Router

	router.HandleFunc("/api/check-harga", app.CheckPrice).Methods(http.MethodGet)

	log.Printf("api running in port %d", c.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", c.Port), router)
}
