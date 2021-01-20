package nature_remo_sdk

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type NatureRemoSdk struct {
	Token string
}

func (s *NatureRemoSdk) request(method string, url string) ([]byte, error) {
	client := &http.Client{}
	client.Timeout = time.Second * 30

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", s.Token)
	req.Header.Add("accept", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("Status NG: " + strconv.Itoa(res.StatusCode))
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
