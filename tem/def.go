package tem

import (
	"strconv"

	"fmt"

	"github.com/go-chat-bot/bot"
)

func defProject(cmd *bot.Cmd) (string, error) {
	if len(cmd.Args) == 0 {
		return "set default project with index,like `tem defp 0`", nil
	}
	i, err := strconv.Atoi(cmd.Args[0])
	if err != nil {
		return "set default project err:" + err.Error(), nil
	}
	if u, ok := users[cmd.User.ID]; ok {
		if i < 0 || i >= len(u.Projects) {
			return fmt.Sprintf("set default project error:%d/%d", i, len(u.Projects)), nil
		}
		u.DefProjectIndex = i
		return "set default project success,default project:" + u.Projects[i].Name, nil
	}
	return "login first", nil
}

func init() {
	bot.RegisterCommand("defp",
		"default project index",
		"defp 0",
		defProject,
	)
}
