package tem

import (
	"github.com/go-chat-bot/bot"
)

func ls(cmd *bot.Cmd) (string, error) {
	if u, ok := users[cmd.User.ID]; ok {
		return u.GetProject()
	}
	return "login first", nil
}

func init() {
	bot.RegisterCommand("ls",
		"list something",
		"ls",
		ls,
	)
}
