package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	tb "gopkg.in/telebot.v3"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Помилка при завантаженні .env")
	}

	token := os.Getenv("TELE_TOKEN")
	if token == "" {
		log.Fatal("TELE_TOKEN не заданий")
	}

	pref := tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := tb.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	bot.Handle(tb.OnText, func(c tb.Context) error {
		return c.Send("Привіт! Я отримав твоє повідомлення: " + c.Text())
	})

	log.Println("Бот запущено!")
	bot.Start()
}
