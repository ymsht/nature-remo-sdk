package nature_remo_sdk

import (
	"encoding/json"
	"time"
)

type Device struct {
	ID                 string    `json:"id"`
	Name               string    `json:"name"`
	Temperature_offset int       `json:"temperature_offset"`
	Humidity_offset    int       `json:"humidity_offset"`
	Created_at         time.Time `json:"created_at"`
	Updated_at         time.Time `json:"updated_at"`
	Firmware_version   string    `json:"firmware_version"`
	Mac_address        string    `json:"mac_address"`
	Serial_number      string    `json:"serial_number"`
	Newest_events      struct {
		Te struct {
			Val        int       `json:"val"`
			Created_at time.Time `json:"created_at"`
		} `json:"newest_events"`
		Hu struct {
			Val        int       `json:"val"`
			Created_at time.Time `json:"created_at"`
		} `json:"hu"`
		Il struct {
			Val        int       `json:"val"`
			Created_at time.Time `json:"created_at"`
		} `json:"il"`
		Mo struct {
			Val        int       `json:"val"`
			Created_at time.Time `json:"created_at"`
		} `json:"mo"`
	} `json:"newest_events"`
}

const url_devices = "https://api.nature.global/1/devices"

// TODO 作成中
func (s *NatureRemoSdk) GetDevice() (Device, error) {
	var device Device

	body, err := s.request("GET", url_devices)
	if err != nil {
		return device, err
	}

	err = json.Unmarshal(body, &device)
	if err != nil {
		return device, err
	}

	return device, nil
}
