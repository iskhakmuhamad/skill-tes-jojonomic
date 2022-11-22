package main

import (
	"fmt"
	"log"
	"net/http"

	"mutation-check-service/configs"
	"mutation-check-service/internal/app"
)

func main() {
	c := configs.NewConfig()
	router := c.Router

	router.HandleFunc("/api/mutasi", app.CheckMutation).Methods(http.MethodGet)

	log.Printf("api running in port %d", c.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", c.Port), router)
}
