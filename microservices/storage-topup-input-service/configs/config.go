package configs

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"storage-topup-input-service/internal/models"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	kafka "github.com/segmentio/kafka-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Config struct {
	Router *mux.Router
	Kafka  *kafka.Reader
	DB     *gorm.DB
	Port   int
}

func NewConfig() *Config {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	rt := mux.NewRouter()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.Topup{})
	db.AutoMigrate(&models.Account{})
	db.AutoMigrate(&models.Transaction{})

	brokers := strings.Split(os.Getenv("KAFKA_URL"), ",")
	kafka := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  os.Getenv("KAFKA_GROUP_ID"),
		Topic:    os.Getenv("KAFKA_TOPIC"),
		MinBytes: 0,    // 1KB
		MaxBytes: 10e6, // 10MB
	})

	port := os.Getenv("PORT")
	portInt, _ := strconv.Atoi(port)

	app := &Config{
		Router: rt,
		Kafka:  kafka,
		Port:   portInt,
		DB:     db,
	}

	return app
}
