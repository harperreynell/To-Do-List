package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Client() {
	bot, err := tgbotapi.NewBotAPI("")
	if err != nil {
		panic(err)
	}

	bot.Debug = true
}
