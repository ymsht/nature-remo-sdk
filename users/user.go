package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
}

func GetMe() {
	url := "https://api.nature.global/1/users/me"
	token := ""
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Bearer", token)
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

	fmt.Println(string(body))

	var users []User
	err = json.Unmarshal(body, &users)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(users); i++ {
		fmt.Println(users[i])
	}
}
