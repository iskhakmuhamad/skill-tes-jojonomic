package configs

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"mutation-check-service/internal/models"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Config struct {
	Router *mux.Router
	DB     *gorm.DB
	Port   int
}

func NewConfig() *Config {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	rt := mux.NewRouter()

	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

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
		Logger: dbLogger,
	})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.Transaction{})

	port := os.Getenv("PORT")
	portInt, _ := strconv.Atoi(port)

	app := &Config{
		Router: rt,
		Port:   portInt,
		DB:     db,
	}

	return app
}
