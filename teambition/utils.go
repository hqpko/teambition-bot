package teambition

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func request(url, token string) ([]byte, error) {
	fmt.Println("request:", url)
	req, err := getRequest(url, token)
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

func getRequest(url, token string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "OAuth2 "+token)
	return req, nil
}
