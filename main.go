package main

import (
	"os"

	"github.com/go-chat-bot/bot"
	"github.com/go-chat-bot/bot/slack"
	"github.com/hqpko/teambition-bot/tem"
)

func main() {
	bot.SetBotName("tem")
	tem.SetTeambitionAPIToken(os.Getenv("TEAMBITION_API_TOKEN"))
	slack.Run(os.Getenv("SLACK_TOKEN"))
}
