package app

import (
	"context"
	"encoding/json"
	"log"
	"storage-topup-input-service/configs"
	"storage-topup-input-service/internal/models"
	"strconv"

	"github.com/teris-io/shortid"
	"gorm.io/gorm"
)

func ReadMessage() {
	c := configs.NewConfig()
	ctx := context.Background()

	for {
		kf, err := c.Kafka.FetchMessage(ctx)
		if err != nil {
			break
		}
		log.Printf("message at topic/partition/offset %v/%v/%v: %s\n", kf.Topic, kf.Partition, kf.Offset, string(kf.Key))

		m := new(models.Topup)
		if err := json.Unmarshal(kf.Value, &m); err != nil {
			log.Printf("unmarshall data error : %s", err.Error())
		}

		if err := c.DB.Model(m).Create(m).Error; err != nil {
			log.Printf("error when insert topup to database with error message %s\n", err.Error())
		}

		saveToAccount(c.DB, m)
		saveToTransaction(c.DB, m)

		if err := c.Kafka.CommitMessages(ctx, kf); err != nil {
			log.Printf("CommitMessage Failed error : %s", err.Error())
		}
	}
}

func saveToAccount(db *gorm.DB, data *models.Topup) {
	account := getAccount(db, data.AccountNo)

	gram, err := strconv.ParseFloat(data.Gram, 64)
	if err != nil {
		log.Printf("error when parse price with error message %s\n", err.Error())
	}

	if account.ReffID != `` {
		balance := account.Balance + gram
		account.Balance = balance

		if err := db.Model(account).Where("reff_id = ?", account.ReffID).Updates(&account).Error; err != nil {
			log.Printf("error when insert account to database with error message %s\n", err.Error())
		}
	}

	if account.ReffID == `` {
		id, _ := shortid.Generate()
		r := &models.Account{
			ReffID: id,
			AccountData: models.AccountData{
				AccountNo: data.AccountNo,
				Balance:   gram,
			},
		}

		if err := db.Model(r).Save(&r).Error; err != nil {
			log.Printf("error when insert account to database with error message %s\n", err.Error())
		}
	}
}

func saveToTransaction(db *gorm.DB, data *models.Topup) {
	account := getAccount(db, data.AccountNo)
	price := getCurrentPrice(db)

	gram, err := strconv.ParseFloat(data.Gram, 64)
	if err != nil {
		log.Printf("error when parse price with error message %s\n", err.Error())
	}

	id, _ := shortid.Generate()
	m := &models.Transaction{
		ReffID: id,
		TransactionData: models.TransactionData{
			Type:         "topup",
			Gram:         gram,
			Balance:      account.Balance,
			PriceTopup:   price.PriceTopup,
			PriceBuyback: price.PriceBuyback,
			AccountNo:    data.AccountNo,
		},
	}

	if err := db.Model(m).Save(&m).Error; err != nil {
		log.Printf("error when insert transaction to database with error message %s\n", err.Error())
	}
}

func getAccount(db *gorm.DB, accountNo string) *models.Account {
	var account *models.Account
	if err := db.Model(account).Where("account_no = ?", accountNo).First(&account).Error; err != nil {
		log.Printf("error when get data Account from database with error message %s\n", err.Error())
	}

	return account
}

func getCurrentPrice(db *gorm.DB) *models.Price {
	var h *models.Price
	if err := db.Model(h).Order("created_at DESC").First(&h).Error; err != nil {
		log.Printf("error when get data price from database with error message %s\n", err.Error())
	}

	return h
}
