package picnic

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	ApiKey string
	Url    string
}

func (cl Client) get(endpoint string, payload ...[]byte) ([]byte, error) {
	client := &http.Client{}
	var body []byte
	if len(payload) > 1 {
		fmt.Println("You should only use one or no variable for the payload. Assigning first value to payload for now.")
		body = payload[0]
	} else if len(payload) > 0 {
		body = payload[0]
	} else {
		body = nil
	}
	req, _ := http.NewRequest("GET", cl.Url+endpoint, bytes.NewBuffer([]byte(body)))

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

func (cl Client) post(endpoint string, payload ...[]byte) ([]byte, error) {
	client := &http.Client{}
	var body []byte
	if len(payload) > 1 {
		fmt.Println("You should only use one or no variable for the payload. Assigning first value to payload for now.")
		body = payload[0]
	} else if len(payload) > 0 {
		body = payload[0]
	} else {
		body = nil
	}
	req, _ := http.NewRequest("POST", cl.Url+endpoint, bytes.NewBuffer([]byte(body)))
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
