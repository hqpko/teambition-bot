package teambition

import (
	"encoding/json"
	"fmt"
)

const (
	teambitionAPIURL = "https://api.teambition.com/"
)

//User teambition 用户数据，只保留用到的数据，其他的忽略
type User struct {
	ID    string `json:"_id,omitempty"`
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
	Phone string `json:"phone,omitempty"`

	Projects         []*Project `json:"-"`
	DefProjectIndex  int        `json:"-"`
	DefTaskListIndex int        `json:"-"`

	Token string `json:"-"`
}

//InitInfo _
func (u *User) InitInfo() error {
	data, err := request(teambitionAPIURL+"api/users/me", u.Token)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, u)
}

func (u *User) UpdateProject() error {
	data, err := request(teambitionAPIURL+"api/projects", u.Token)
	if err != nil {
		return err
	}
	o := []*Project{}
	err = json.Unmarshal(data, &o)
	if err != nil {
		return err
	}
	u.Projects = o
	return nil
}

func (u *User) GetProject() (string, error) {
	s := ""
	for i := 0; i < len(u.Projects); i++ {
		if s != "" {
			s += "\n"
		}
		s += fmt.Sprintf("%d %s", i, u.Projects[i].Name)
	}
	return s, nil
}

func (u *User) GetTaskLists() (string, error) {
	defProject := u.Projects[u.DefProjectIndex]
	s := ""
	for i := 0; i < len(defProject.TaskLists); i++ {
		if s != "" {
			s += "\n"
		}
		s += fmt.Sprintf("%d %s", i, defProject.TaskLists[i].Title)
	}
	return s, nil
}

//Update update projects && tasklist
func (u *User) Update() error {
	e := u.UpdateProject()
	if e != nil {
		return e
	}
	for _, p := range u.Projects {
		e := p.updateTaskLists(u.Token)
		if e != nil {
			return e
		}
	}
	return nil
}

func (u *User) GetTasksWithUserID(userID string) (string, error) {
	defProject := u.Projects[u.DefProjectIndex]
	defTaskList := defProject.TaskLists[u.DefTaskListIndex]
	err := defTaskList.updateTasks(u.Token)
	if err != nil {
		return "", err
	}

	s := ""
	for _, t := range defTaskList.Tasks {
		if t.ExecutorId == userID {
			if s != "" {
				s += "\n"
			}
			s += fmt.Sprintf("%s %s %s", u.Name, t.Stage.Name, t.Content)
		}
	}
	return s, nil
}
