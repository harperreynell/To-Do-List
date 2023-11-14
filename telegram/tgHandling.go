package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"strings"
	l "todolist/listHandling"
)

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
	if txt == "" {
		todoList := l.ReadJsonFromFile(message.Chat.ID)
		todoList = append(todoList, l.NewTodo(text, len(todoList)))
		l.WriteJsonToFile(todoList, message.Chat.ID)
		txt = text
		bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Task was added successfully!"))
		sendMenu(message.Chat.ID)
	}

	if idCheck == -1 {
		todoList := l.ReadJsonFromFile(message.Chat.ID)
		index, _ := strconv.Atoi(text)
		id := l.TaskId(todoList, index)
		todoList = l.DeleteTaskByID(todoList, id)
		l.WriteJsonToFile(todoList, message.Chat.ID)
		bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Task was deleted successfully!"))
		idCheck = 0
		sendMenu(message.Chat.ID)
	}
	if idCheckStat == -1 {
		todoList := l.ReadJsonFromFile(message.Chat.ID)
		index, _ := strconv.Atoi(text)
		id := l.TaskId(todoList, index)
		l.ChangeStatus(&todoList[id], "closed")
		l.WriteJsonToFile(todoList, message.Chat.ID)
		bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Task status was changed successfully!"))
		idCheckStat = 0
		sendMenu(message.Chat.ID)
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
	case "/start":
		sendMenu(chatId)
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
		todoList := l.ReadJsonFromFile(message.Chat.ID)
		text = "Your todo list:\n\n" + l.PrintTodos(todoList)
		markup = MenuMarkup
	} else if query.Data == addButton {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Type in task:")
		bot.Send(msg)
		txt = ""
	} else if query.Data == deleteButton {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Type in id of task you want to delete:")
		bot.Send(msg)
		idCheck = -1
	} else if query.Data == changeButton {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Type in id of task you want to change:")
		bot.Send(msg)
		idCheckStat = -1
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
