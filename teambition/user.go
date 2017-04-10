package teambition

import (
	"encoding/json"
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

	Token string `json:"-"`
}

//InitInfo _
func (u *User) InitInfo() error {
	req, err := http.NewRequest("GET", teambitionAPIURL+"api/users/me", nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "OAuth2 "+u.Token)
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, u)
}
