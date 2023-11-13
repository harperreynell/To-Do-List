package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
	l "todolist/listHandling"
)

var (
	// Menu texts
	firstMenu   = "<b>Menu for task list</b>\n\n"
	printButton = "Print tasks"
	addButton   = "Add task"
	txt         = "1"

	bot *tgbotapi.BotAPI

	userState map[int64]string

	MenuMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(printButton, printButton),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(addButton, addButton),
		),
	)
)

func Run() {
	var err error
	bot, err = tgbotapi.NewBotAPI("")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false
	u := tgbotapi.NewUpdate(0)
	updates := bot.GetUpdatesChan(u)
	u.Timeout = 15

	receiveUpdates(updates)
}

func receiveUpdates(updates tgbotapi.UpdatesChannel) {
	for {
		for update := range updates {
			handleUpdate(update)
		}
	}
}

func handleUpdate(update tgbotapi.Update) {
	switch {
	case update.Message != nil:
		handleMessage(update.Message)
		break

	case update.CallbackQuery != nil:
		handleButton(update.CallbackQuery)
		break
	}
}

func handleMessage(message *tgbotapi.Message) {
	user := message.From
	text := message.Text

	if user == nil {
		return
	}

	log.Printf("%s wrote %s", user.FirstName, text)
	log.Println(txt)
	if txt == "" {
		log.Println(txt, text)
		todoList := l.ReadJsonFromFile()
		todoList = append(todoList, l.NewTodo(text, len(todoList)))
		l.WriteJsonToFile(todoList)
		txt = text
	}
	var err error
	if strings.HasPrefix(text, "/") {
		handleCommand(message.Chat.ID, text)
	}

	if err != nil {
		log.Printf("An error occured: %s", err.Error())
	}

}

func handleCommand(chatId int64, command string) {

	switch command {
	case "/menu":
		sendMenu(chatId)
		break
	}
}

func handleButton(query *tgbotapi.CallbackQuery) {
	var text string

	markup := tgbotapi.NewInlineKeyboardMarkup()
	message := query.Message

	if query.Data == printButton {
		todoList := l.ReadJsonFromFile()
		text = "Your todo list:\n\n" + l.PrintTodos(todoList)
		markup = MenuMarkup
	} else if query.Data == addButton {

		msg := tgbotapi.NewMessage(message.Chat.ID, "Type in task:")
		bot.Send(msg)
		txt = ""
	}

	callbackCfg := tgbotapi.NewCallback(query.ID, "")
	bot.Send(callbackCfg)

	msg := tgbotapi.NewEditMessageTextAndMarkup(message.Chat.ID, message.MessageID, text, markup)
	msg.ParseMode = tgbotapi.ModeHTML
	bot.Send(msg)
}

func sendMenu(chatId int64) error {
	msg := tgbotapi.NewMessage(chatId, firstMenu)
	msg.ParseMode = tgbotapi.ModeHTML
	msg.ReplyMarkup = MenuMarkup
	_, err := bot.Send(msg)
	return err
}
