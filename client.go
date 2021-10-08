package picnic

import (
	"io/ioutil"
	"net/http"
)

type Client struct {
	ApiKey string
	Url    string
}

func (cl Client) get(endpoint string) ([]byte, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", cl.Url+endpoint, nil)
	req.Header.Set("x-picnic-auth", cl.ApiKey)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return result, nil
}
