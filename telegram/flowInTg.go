package telegram

//
//import (
//	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
//)
//
//var (
//	updateConfig = tgbotapi.NewUpdate(0)
//	updates      = bot.GetUpdatesChan(updateConfig)
//)

//func Client() {
//
//	updateConfig.Timeout = 10
//
//	for update := range updates {
//		if update.Message == nil {
//			continue
//		}
//		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
//		if strings.HasPrefix(update.Message.Text, "/") {
//			handleCommand(update.Message.Chat.ID, update.Message.Text)
//		}
//
//	}
//	//bot.Debug = true
//}
