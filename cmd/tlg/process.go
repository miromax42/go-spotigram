package tlg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	mainMenuKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("ADD"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("CREATE"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("HELP"),
		),
	)
	addKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("add"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("show"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("back"),
		),
	)
)

func process(update *tgbotapi.Update) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

	if update.Message.IsCommand() {
		switch update.Message.Command() {
		case "start":
			msg.ReplyMarkup = mainMenuKeyboard
		case "close":
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		case "show":
			//todo
		case "add":
			//todo
		case "url":
			msg.ParseMode = "html"
			msg.Text = "This will be interpreted as HTML, click <a href=\"https://www.example.com\">here</a>"
		default:
			msg.Text = "I don't know that command"
		}
	}
	switch update.Message.Text {
	case "ADD":
		msg.ReplyMarkup = addKeyboard
	case "CREATE":
		msg.Text = "//todo"
	case "HELP":
		msg.Text = "This will be help"
	case "add":
		msg.Text = "//todo"
	case "show":
		msg.Text = "//todo"
	case "back":
		msg.ReplyMarkup = mainMenuKeyboard
	}
	return msg
}
