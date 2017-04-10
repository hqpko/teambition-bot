package tem

import "github.com/go-chat-bot/bot"

func login(cmd *bot.Cmd) (string, error) {
	return "login with this:" + teambitionAPI.OAuthURL(cmd.User.ID), nil
}

func init() {
	bot.RegisterCommand(
		"login",
		"teambition login",
		"login",
		login)
}
