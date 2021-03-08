package main

import (
	"github.com/cachemem/GoTeleMeido/internal/reverse"
	"math/rand"
)

const helpText = "Use /help to see available commands."

type AnswerBot struct {
	r reverse.Reverser
}

func NewBot(reverser reverse.Reverser) *AnswerBot {
	if reverser == nil {
		return &AnswerBot{r: reverse.R{}}
	}
	return &AnswerBot{r: reverser}
}

// commands itself
func (ab *AnswerBot) greetings() string {
	return "Hai~ " + helpText
}

func (ab *AnswerBot) unknowCommand() string {
	return "Unknown command. " + helpText
}

func (ab *AnswerBot) eightBall() string {
	answers := []string{
		"It is certain",
		"It is decidedly so",
		"Without a doubt",
		"Yes definitely",
		"You may rely on it",
		"As I see it yes",
		"Most likely",
		"Outlook good",
		"Yes",
		"Signs point to yes",
		"Reply hazy try again",
		"Ask again later",
		"Better not tell you now",
		"Cannot predict now",
		"Concentrate and ask again",
		"Don't count on it",
		"My reply is no",
		"My sources say no",
		"Outlook not so good",
		"Very doubtful",
	}
	return answers[rand.Intn(len(answers))]
}

func (ab *AnswerBot) help() string {
	return "* /reverse — reverse whatever text want\n* /8ball — ask a magic 8-ball"
}

func (ab *AnswerBot) reverse(s *string) string {
	if s == nil {
		return ""
	}

	return ab.r.Reverse(*s)
}

// Main method
func (ab *AnswerBot) ProcessCommand(command string, cmdArgs *string) string {
	switch command {
	case "help":
		return ab.help()
	case "reverse":
		return ab.reverse(cmdArgs)
	case "8ball":
		return ab.eightBall()
	case "start":
		return ab.greetings()
	default:
		return ab.unknowCommand()
	}
}
