package main

import (
	"ReminderAssets/clients/telegram_client"
	"flag"
	"log"
)

func main() {
	tgBotHost := "api.telegram.org"
	tgClient := telegram_client.NewClient(tgBotHost, mustToken())
}

func mustToken() string {
	token := flag.String("token-bot", "", "Токен для запуска бота")
	flag.Parse()

	if *token == "" {
		log.Fatal("Токен не указан")
	}

	return *token
}
