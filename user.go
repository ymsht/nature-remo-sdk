package nature_remo_sdk

import (
	"encoding/json"
)

type User struct {
	ID       string `json:"id"`
	Nickname string `json:"nickname"`
}

const url = "https://api.nature.global/1/users/me"

// TODO 作成中
func (s *NatureRemoSdk) GetMe() (User, error) {
	var user User

	body, err := s.request("GET", url)
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		return user, err
	}

	return user, nil
}
