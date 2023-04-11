package main

import (
	"ReminderAssets/clients/telegram_client"
	"flag"
	"fmt"
	"log"
)

func main() {
	tgBotHost := "api.telegram.org"
	tgClient := telegram_client.NewClient(tgBotHost, mustToken())
	fmt.Println(tgClient)
}

func mustToken() string {
	token := flag.String("token-bot", "", "Токен для запуска бота")
	flag.Parse()

	if *token == "" {
		log.Fatal("Токен не указан")
	}

	return *token
}
