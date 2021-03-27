package main

import (
	"github.com/cachemem/GoTeleMeido/internal/repository/postgres"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	botToken := os.Getenv("TOKEN")
	dbUrl := os.Getenv("DB_URL")

	database := postgres.New(dbUrl)
	defer database.Close()

	botTransport, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic("Can not initialize bot")
	}
	botTransport.Debug = true
	log.Printf("Authorized on account %s", botTransport.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := botTransport.GetUpdatesChan(u)

	bot := NewBot(nil, database)
	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		if update.Message.IsCommand() {
			cmd := update.Message.Command()
			cmdArgs := update.Message.CommandArguments()
			msg.Text = bot.ProcessCommand(cmd, &cmdArgs, int64(update.Message.From.ID))
		} else {
			msg.Text = bot.ProcessCommand("", nil, int64(update.Message.From.ID))
		}

		botTransport.Send(msg)
	}
}
