package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var (
	// Menu texts
	firstMenu    = "<b>Menu for task list</b>\n\n"
	printButton  = "Print tasks"
	addButton    = "Add task"
	deleteButton = "Delete task"
	changeButton = "Change task status"
	txt          = "1"
	idCheck      = 0
	stat         = "-"
	idCheckStat  = 0

	bot *tgbotapi.BotAPI

	MenuMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(printButton, printButton),
			tgbotapi.NewInlineKeyboardButtonData(addButton, addButton),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(deleteButton, deleteButton),
			tgbotapi.NewInlineKeyboardButtonData(changeButton, changeButton),
		),
	)
)
