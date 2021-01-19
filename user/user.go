package user

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type User struct {
	ID       string `json:"id"`
	Nickname string `json:"nickname"`
}

func GetMe() User {
	// TODO 作成中
	url := "https://api.nature.global/1/users/me"
	token := "Bearer "
	client := &http.Client{}
	client.Timeout = time.Second * 30

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", token)
	req.Header.Add("accept", "application/json")
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal(res)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Fatal(err)
	}

	return user
}
