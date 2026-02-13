package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"main/game"
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

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			if update.Message.Text == "/start" {
				// Создаем клавиатуру с кнопкой "Начать игру"
				keyboard := tgbotapi.NewInlineKeyboardMarkup(
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonData("Начать игру", "start_game"),
					),
				)

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Нажми кнопку ниже, чтобы начать игру:")
				msg.ReplyMarkup = keyboard

				bot.Send(msg)
				continue
			}

			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			word := game.Word(update.Message.Text)
			if word == "ты проиграл!" {
				// Создаем клавиатуру с кнопкой "Перезапустить игру"
				keyboard := tgbotapi.NewInlineKeyboardMarkup(
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonData("Перезапустить игру", "restart_game"),
					),
				)

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ты проиграл!! Нажми кнопку ниже, чтобы перезапустить игру:")
				msg.ReplyMarkup = keyboard

				bot.Send(msg)
				continue
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, word)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		} else if update.CallbackQuery != nil {
			// Обрабатываем нажатие кнопки
			if update.CallbackQuery.Data == "start_game" {
				// Отправляем сообщение о начале игры
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Игра началась! Введи первое слово:")
				bot.Send(msg)
			}
			if update.CallbackQuery.Data == "restart_game" {
				// Отправляем сообщение о начале игры
				game.RestartGame()
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Игра началась! Введи первое слово:")
				bot.Send(msg)
			}
		}
	}
}
