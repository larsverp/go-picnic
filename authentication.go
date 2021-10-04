package picnic

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net/http"
)

type User struct {
	Key       string `json:"key"`
	Secret    string `json:"secret"`
	Client_id int    `json:"client_id"`
}

type Client struct {
	ApiKey string
	url    string
}

const url string = "https://storefront-prod.nl.picnicinternational.com/api/15"

func NewUser(email string, password string) User {
	user := User{Key: email, Secret: getMD5Hash(password), Client_id: 1}
	return user
}

func NewClient(user User) (*Client, error) {
	jsonData, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(url+"/user/login", "application/json", bytes.NewBuffer(jsonData))
	if err != nil || resp.StatusCode != 200 {
		return nil, err
	}

	client := Client{ApiKey: resp.Header["X-Picnic-Auth"][0], url: url}

	return &client, nil
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
