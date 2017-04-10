package tem

import (
	"github.com/go-chat-bot/bot"
)

func lp(cmd *bot.Cmd) (string, error) {
	if u, ok := users[cmd.User.ID]; ok {
		return u.GetProject()
	}
	return "login first", nil
}

func lt(cmd *bot.Cmd) (string, error) {
	if u, ok := users[cmd.User.ID]; ok {
		return u.GetTaskLists()
	}
	return "login first", nil
}

func ls(cmd *bot.Cmd) (string, error) {
	if u, ok := users[cmd.User.ID]; ok {
		return u.GetTasksWithUserID(u.ID)
	}
	return "login first", nil
}

func init() {
	bot.RegisterCommand("lp",
		"list something",
		"lp",
		lp,
	)
	bot.RegisterCommand("lt",
		"list tasklist",
		"lt",
		lt,
	)
	bot.RegisterCommand("ls",
		"list tasks",
		"ls",
		ls,
	)
}
