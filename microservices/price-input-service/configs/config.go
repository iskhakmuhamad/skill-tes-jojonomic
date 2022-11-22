package configs

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	kafka "github.com/segmentio/kafka-go"
)

type Config struct {
	Router *mux.Router
	Kafka  *kafka.Conn
	Port   int
}

func NewConfig() *Config {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	rt := mux.NewRouter()

	kafka, err := kafka.DialLeader(context.Background(), "tcp", os.Getenv("KAFKA_URL"), os.Getenv("KAFKA_TOPIC"), 0)
	if err != nil {
		log.Fatalf("kafka connection err : %v", err.Error())
	}

	port := os.Getenv("PORT")
	portInt, _ := strconv.Atoi(port)

	app := &Config{
		Router: rt,
		Kafka:  kafka,
		Port:   portInt,
	}

	return app
}
