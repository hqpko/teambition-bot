package teambition

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
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
	data, err := u.request(teambitionAPIURL + "api/users/me")
	if err != nil {
		return err
	}
	return json.Unmarshal(data, u)
}

func (u *User) UpdateProject() error {
	data, err := u.request(teambitionAPIURL + "api/projects")
	if err != nil {
		return err
	}
	o := []*Project{}
	err = json.Unmarshal(data, &o)
	if err != nil {
		return err
	}
	u.Projects = o
	fmt.Println(string(data))
	fmt.Println(o)
	return nil
}

func (u *User) UpdateTaskLists() error {
	if len(u.Projects) <= u.DefProjectIndex {
		return errors.New("projects is empty.")
	}
	defProject := u.Projects[u.DefProjectIndex]
	data, err := u.request(teambitionAPIURL + "api/projects/" + defProject.ID + "/tasklists")
	if err != nil {
		return err
	}
	o := []*TaskList{}
	err = json.Unmarshal(data, &o)
	if err != nil {
		return err
	}
	defProject.TaskLists = o
	fmt.Println(string(data))
	fmt.Println(o)
	return nil
}

func (u *User) GetProject() (string, error) {
	if err := u.UpdateProject(); err != nil {
		return "", err
	}
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
	if err := u.UpdateTaskLists(); err != nil {
		return "", err
	}
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

func (u *User) request(url string) ([]byte, error) {
	fmt.Println("request:", url)
	req, err := u.getRequest(url)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (u *User) getRequest(url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "OAuth2 "+u.Token)
	return req, nil
}
