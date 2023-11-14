package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

func Run() {
	var err error
	bot, err = tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false
	u := tgbotapi.NewUpdate(0)
	updates := bot.GetUpdatesChan(u)
	u.Timeout = 15

	receiveUpdates(updates)
}
