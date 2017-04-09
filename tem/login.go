package tem

import (
	"net/url"

	"fmt"

	"github.com/go-chat-bot/bot"
)

func login(cmd *bot.Cmd) (string, error) {
	values := url.Values{}
	values.Set("client_id", teambitionAPIToken)
	values.Set("response_type", "code")
	values.Set("redirect_uri", fmt.Sprintf("http://%s:%s/login", temHost, temPort))
	values.Set("state", cmd.User.ID)
	return "https://account.teambition.com/oauth2/authorize?" + values.Encode(), nil
}

func init() {
	bot.RegisterCommand(
		"login",
		"teambition login",
		"login",
		login)
}
