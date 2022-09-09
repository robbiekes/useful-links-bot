package main

import (
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/robbiekes/LinksBot/linksbot/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	storage := initStorage()
	fmt.Println(storage.PickRandom("user1"))
	// repository.Save(&repository.Page{"link1", "user1"})
	// fmt.Println(repository.IsPresent(&repository.Page{"aa", "user"}))

	// token := mustToken()
	// host := mustHost()
	// tgClient := telegram.NewClient(host, token)

	// fetcher = fetcher.New(tgClient)
	// processor = processor.New(tgClient)

	// consumer.Start(fetcher, processor)
}

func initStorage() *repository.StoragePostgres {
	if err := initConfig(); err != nil {
		logrus.Fatalf("failed to connect to database: %s", err)
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("failed to connect to database: %s", err)
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("host"),
		Port:     viper.GetString("port"),
		Username: viper.GetString("user"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("dbname"),
		SSlMode:  viper.GetString("sslmode"),
	})
	if err != nil {
		log.Fatalf("failed to connect to database: %s", err)
	}

	storage := repository.NewStoragePostgres(db)
	return storage
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func mustToken() string { // must позволяет не возвращать ошибку и обрабатывать её внутри функции
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN is not set")
	}
	return token
}

func mustHost() string {
	host := os.Getenv("TELEGRAM_BOT_HOST")
	if host == "" {
		log.Fatal("TELEGRAM_BOT_HOST is not set")
	}
	return host
}
