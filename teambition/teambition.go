package teambition

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	teambitionOAuthURL       = "https://account.teambition.com/oauth2/authorize"
	teambitionAccessTokenURL = "https://account.teambition.com/oauth2/access_token"
)

//TeambitionAPI 主要用于认证
type TeambitionAPI struct {
	appKey      string //申请应用时分配的 AppKey
	appSecret   string //申请应用时分配的 AppSecret
	redirectURL string //授权回调地址，站外应用需与设置的回调地址一致
}

//CreateTeambitionAPI _
func CreateTeambitionAPI(appKey, appSecret, redirectURL string) *TeambitionAPI {
	return &TeambitionAPI{appKey: appKey, appSecret: appSecret, redirectURL: redirectURL}
}

//OAuthURL 获取认证地址
//param state string 用于回调的返回参数
func (t *TeambitionAPI) OAuthURL(state string) string {
	values := url.Values{}
	values.Set("client_id", t.appKey)
	values.Set("redirect_uri", t.redirectURL)
	values.Set("response_type", "code")
	values.Set("state", state)
	return teambitionOAuthURL + "?" + values.Encode()
}

//GetAccessToken 获取用户的token
//param code string oauth登录的授权码
func (t *TeambitionAPI) getAccessToken(code string) (string, error) {
	values := url.Values{}
	values.Set("client_id", t.appKey)
	values.Set("client_secret", t.appSecret)
	values.Set("code", code)
	resp, err := http.PostForm(teambitionAccessTokenURL, values)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	token := &struct {
		AccessToken string `json:"access_token"`
	}{}
	err = json.Unmarshal(data, token)
	if err != nil {
		return "", err
	}
	return token.AccessToken, nil
}

//Login oauth方式登录，并返回用户数据
func (t *TeambitionAPI) Login(code string) (*User, error) {
	token, err := t.getAccessToken(code)
	if err != nil {
		return nil, err
	}
	u := &User{Token: token}
	err = u.InitInfo()
	if err != nil {
		return nil, err
	}
	err = u.Update()
	if err != nil {
		return nil, err
	}
	return u, nil
}
