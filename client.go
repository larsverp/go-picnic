package picnic

import (
	"bytes"
	"errors"
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
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("API did not return an OK status code, check body for more info")
	}

	return result, nil
}

func (cl Client) post(endpoint string, payload []byte) ([]byte, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", cl.Url+endpoint, bytes.NewBuffer([]byte(payload)))
	req.Header.Set("x-picnic-auth", cl.ApiKey)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("API did not return an OK status code, check body for more info")
	}

	return result, nil
}
