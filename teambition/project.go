package teambition

import "encoding/json"

type Project struct {
	ID   string `json:"_id,omitempty"`
	Name string `json:"name,omitempty"`

	TaskLists []*TaskList `json:"-"`
}

func (p *Project) updateTaskLists(token string) error {
	data, err := request(teambitionAPIURL+"api/projects/"+p.ID+"/tasklists", token)
	if err != nil {
		return err
	}
	o := []*TaskList{}
	err = json.Unmarshal(data, &o)
	if err != nil {
		return err
	}
	p.TaskLists = o
	return nil
}
