package main

import (
	"fmt"
	"log"
	"net/http"

	"balance-check-service/configs"
	"balance-check-service/internal/app"
)

func main() {
	c := configs.NewConfig()
	router := c.Router

	router.HandleFunc("/api/saldo", app.BalanceCheck).Methods(http.MethodGet)

	log.Printf("api running in port %d", c.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", c.Port), router)
}
