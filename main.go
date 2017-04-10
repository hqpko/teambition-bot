package main

import (
	"os"

	"github.com/go-chat-bot/bot"
	"github.com/go-chat-bot/bot/slack"
	"github.com/hqpko/teambition-bot/tem"
)

func main() {
	bot.SetBotName("tem")
	tem.InitTeambitionAPI(os.Getenv("TEAMBITION_APP_KEY"), os.Getenv("TEAMBITION_APP_SECRET"), os.Getenv("REDIRECT_URL"))
	tem.Start(os.Getenv("TEM_HOST"), os.Getenv("TEM_PORT"))
	slack.Run(os.Getenv("SLACK_TOKEN"))
}
