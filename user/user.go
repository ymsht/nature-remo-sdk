package user

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type NatureRemoSdk struct {
	Token string
}

type User struct {
	ID       string `json:"id"`
	Nickname string `json:"nickname"`
}

const url = "https://api.nature.global/1/users/me"

func (s *NatureRemoSdk) GetMe() User {
	// TODO 作成中
	client := &http.Client{}
	client.Timeout = time.Second * 30

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", s.Token)
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
