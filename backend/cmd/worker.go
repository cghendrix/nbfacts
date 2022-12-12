package cmd

import (
	"cghendrix/nbfacts/db"
	"cghendrix/nbfacts/internal/facts"
	"cghendrix/nbfacts/internal/sms"
	"cghendrix/nbfacts/models"
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"log"
)

func setupDB() *sqlx.DB {
	dbUrl := viper.Get("DB_URL").(string)
	dbUser := viper.Get("DB_USER").(string)
	dbPass := viper.Get("DB_PASS").(string)
	dbString := dbUser + ":" + dbPass + "@" + dbUrl
	return db.Init(dbString)
}

func setupServices(db *sqlx.DB) (sms.Service, facts.Service) {
	viper.SetConfigFile("./env/.env")
	viper.ReadInConfig()

	// Sms
	smsRepository := sms.NewRepository(db)
	// facts
	factsRepository := facts.NewRepository(db)
	return sms.NewService(smsRepository), facts.NewService(factsRepository)
}

func setupHandler(smsService sms.Service, factsService facts.Service) sms.Handler {
	return sms.NewHandler(smsService, factsService)
}

func startPubSubWorker() {
	nbDb := setupDB()
	smsService, factsService := setupServices(nbDb)
	handler := setupHandler(smsService, factsService)

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "nicklebackfacts")
	if err != nil {
		log.Fatalf("pubsub.NewClient: %v", err)
	}
	defer client.Close()
	sub := client.Subscription("new-sms-subscription")
	err = sub.Receive(ctx, func(_ context.Context, msg *pubsub.Message) {
		log.Println("new message received")
		var smsMessage models.SMSMessage
		if err := json.Unmarshal(msg.Data, &smsMessage); err != nil {
			panic(err)
		}
		tx, err := nbDb.BeginTx(ctx, nil)
		if err != nil {
			panic(err)
		}
		defer tx.Rollback()
		err = handler.AddSMSWithFact(ctx, tx, sms.CreateSMSRequest{
			Body:       smsMessage.Body,
			From:       smsMessage.From,
			MessageSID: smsMessage.MessageSid,
		})
		if err != nil {
			log.Println("AddSMS error", err)
		}
		// Commit the transaction.
		if err = tx.Commit(); err != nil {
			return
		}
		msg.Ack()
		log.Println("new message ack'd")
	})
	if err != nil {
		fmt.Printf("sub.Receive: %v", err)
	}
}
