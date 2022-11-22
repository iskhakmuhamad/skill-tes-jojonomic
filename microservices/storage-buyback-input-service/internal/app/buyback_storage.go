package app

import (
	"context"
	"encoding/json"
	"log"

	"storage-buyback-input-service/configs"
	"storage-buyback-input-service/internal/models"

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

		m := new(models.BuybackData)
		if err := json.Unmarshal(kf.Value, &m); err != nil {
			log.Printf("unmarshall data error : %s", err.Error())
		}

		if err := c.DB.Model(m).Create(m).Error; err != nil {
			log.Printf("error when insert topup to database with error message %s\n", err.Error())
		}

		go saveToAccount(c.DB, m)
		go saveToTransaction(c.DB, m)

		if err := c.Kafka.CommitMessages(ctx, kf); err != nil {
			log.Printf("CommitMessage Failed error : %s", err.Error())
		}
	}
}

func saveToAccount(db *gorm.DB, data *models.BuybackData) {
	account := getAccount(db, data.AccountNo)

	balance := account.Balance - data.Gram
	account.Balance = balance

	if err := db.Model(account).Where("reff_id = ?", account.ReffID).Updates(&account).Error; err != nil {
		log.Printf("error when insert account to database with error message %s\n", err.Error())
	}
}

func saveToTransaction(db *gorm.DB, data *models.BuybackData) {
	account := getAccount(db, data.AccountNo)

	id, _ := shortid.Generate()
	m := &models.Transaction{
		ReffID: id,
		TransactionData: models.TransactionData{
			Type:         "buyback",
			Gram:         data.Gram,
			Balance:      account.Balance - data.Gram,
			PriceTopup:   data.PriceTopup,
			PriceBuyback: data.PriceBuyback,
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
		log.Printf("error when get data account from database with error message %s\n", err.Error())
	}

	return account
}
