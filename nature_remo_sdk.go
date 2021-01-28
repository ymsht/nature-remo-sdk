package nature_remo_sdk

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"math/rand"
)

type NatureRemoSdk struct {
	Token string
}

func (s NatureRemoSdk) request(method string, url string) ([]byte, error) {
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
		rand.Seed(time.Now().UnixNano())
		upperLimit := 3
		retryNum := 0
		for {
			res, err = client.Do(req)
			if err != nil {
				if retryNum > upperLimit {
                	return nil, err
				}
			
				waitTime := 2 ^ retryNum + rand.Intn(1000) / 1000
				time.Sleep(time.Duration(waitTime) * time.Second)
				retryNum++
				continue
			}
			break
		}
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
