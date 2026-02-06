package main

import (
	"log"
	"main/game"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	game, err := game.NewGame()
	if err != nil {
		log.Fatal(err)
	}

	bot, err := tgbotapi.NewBotAPI("8314606814:AAH7IrFp5jDhaqKMOy-n29y8VUSkP3CvtvU")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			word := game.Word(update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, word)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
