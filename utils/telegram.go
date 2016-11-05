package utils

import (
	"os"
	"strconv"

	"fmt"

	"gopkg.in/telegram-bot-api.v4"
)

func SendTelegramMessage(message string) {

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_KEY"))
	if err != nil {
		fmt.Println(err.Error())
	}

	if bot.Self.UserName == "" {
		fmt.Println("Error connecting to Telegram!")
		return
	}

	chatID, _ := strconv.ParseInt(os.Getenv("TELEGRAM_OWNER_CHATID"), 10, 64)
	msg := tgbotapi.NewMessage(chatID, message)
	bot.Send(msg)

}
