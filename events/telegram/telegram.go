package telegram

import "ReminderAssets/clients/telegram_client"

type Processor struct {
	tg     *telegram_client.Client
	offset int
	//storage
}

func new() {

}
