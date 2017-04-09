package main

import (
	"os"

	"github.com/hqpko/bot"
	"github.com/hqpko/bot/slack"
)

func main() {
	bot.SetBotName("tem")
	slack.Run(os.Getenv("SLACK_TOKEN"))
}
