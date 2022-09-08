package main

import (
	"github.com/robbiekes/LinksBot/client/telegram"
	"log"
	"os"
)

func main() {
	token := mustToken()
	host := mustHost()
	tgClient := telegram.NewClient(host, token)

	// fetcher = fetcher.New(tgClient)
	// processor = processor.New(tgClient)

	// consumer.Start(fetcher, processor)
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
