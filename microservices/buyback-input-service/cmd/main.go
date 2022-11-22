package main

import (
	"fmt"
	"log"
	"net/http"

	"buyback-input-service/configs"
	"buyback-input-service/internal/app"
)

func main() {
	c := configs.NewConfig()
	router := c.Router

	router.HandleFunc("/api/buyback", app.Buyback).Methods(http.MethodPost)

	log.Printf("api running in port %d", c.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", c.Port), router)
}
