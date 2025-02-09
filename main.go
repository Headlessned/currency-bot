package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gopkg.in/telebot.v4"
	"log"
	"os"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("TG_BOT_TOKEN")
	if token == "" {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Токен успешно загружен", token[:10]+"...")

	bot, err := telebot.NewBot(telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}
	bot.Handle("/start", func(c telebot.Context) error {
		return c.Send("Привет! Меня зовут Боб, я бот конвертер валют. " +
			"Напиши /help, чтобы узнать как меня использовать ")
	})

	fmt.Println("Бот Боб запущен бро")
	bot.Start()
}
