package teambition

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type TaskList struct {
	ID    string `json:"_id,omitempty"`
	Title string `json:"title,omitempty"`

	Tasks []*Task `json:"-"`
}

func (t *TaskList) updateTasks(token string) error {
	values := url.Values{}
	values.Set("is_done", "false")
	data, err := request(teambitionAPIURL+"api/tasklists/"+t.ID+"/tasks?"+values.Encode(), token)
	if err != nil {
		return err
	}
	ts := []*Task{}
	err = json.Unmarshal(data, &ts)
	if err != nil {
		return err
	}

	t.Tasks = ts
	fmt.Println(string(data))
	return nil
}
