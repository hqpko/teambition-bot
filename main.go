package main

import (
	"os"

	"github.com/go-chat-bot/bot"
	"github.com/go-chat-bot/bot/slack"
)

func main() {
	bot.SetBotName("tem")
	slack.Run(os.Getenv("SLACK_TOKEN"))
}
