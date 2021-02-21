package main

import (
    "log"
    "os"
    "github.com/cachemem/GoTeleMeido/internal/reverse"
    "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
    botToken := os.Getenv("TOKEN")

    bot, err := tgbotapi.NewBotAPI(botToken)
    if err != nil {
        log.Panic("Can not initialize bot")
    }
    bot.Debug = true
    log.Printf("Authorized on account %s", bot.Self.UserName)

    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60

    updates, err := bot.GetUpdatesChan(u)

    for update := range updates {
        if update.Message == nil { // ignore any non-Message Updates
            continue
        }

        log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

        reversed_response := reverse.Reverse(update.Message.Text)
        msg := tgbotapi.NewMessage(update.Message.Chat.ID, reversed_response)

        bot.Send(msg)
    }
}
