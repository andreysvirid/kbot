package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	token := os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		log.Fatal("No TELEGRAM_TOKEN provided")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// Ініціалізація генератора випадкових чисел
	rand.Seed(time.Now().UnixNano())

	for update := range updates {
		if update.Message == nil {
			continue
		}

		// Генерація випадкового числа від 1 до 6
		roll := rand.Intn(6) + 1

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, 
		                            "🎲 Ви кинули кості: " + 
		                            string(rune('0'+roll))) // перетворення int -> string

		_, err := bot.Send(msg)
		if err != nil {
			log.Println("Failed to send message:", err)
		}
	}
}
