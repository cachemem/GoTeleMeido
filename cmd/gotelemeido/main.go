package main

import (
	"database/sql"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	botToken := os.Getenv("TOKEN")
	dbUrl := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or another initialization error.
		log.Fatal(err)
	}
	log.Printf("Db ping is: %s", db.Ping())
	defer db.Close()

	botTransport, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic("Can not initialize bot")
	}
	botTransport.Debug = true
	log.Printf("Authorized on account %s", botTransport.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := botTransport.GetUpdatesChan(u)

	// Getting ownerId from DB
	var ownerId int64
	rows, err := db.Query("SELECT id FROM users WHERE is_owner = $1 LIMIT 1", true)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&ownerId)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	bot := NewBot(nil, ownerId)
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
