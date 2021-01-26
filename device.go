package nature_remo_sdk

import (
	"encoding/json"
	"time"
)

type sensorValue struct {
	Val        float32   `json:"val"`
	Created_at time.Time `json:"created_at"`
}

type Device struct {
	ID                 string    `json:"id"`
	Name               string    `json:"name"`
	TemperatureOffset int       `json:"temperature_offset"`
	HumidityOffset    int       `json:"humidity_offset"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	FirmwareVersion   string    `json:"firmware_version"`
	MacAddress        string    `json:"mac_address"`
	SerialNumber      string    `json:"serial_number"`
	NewestEvents      struct {
		Te sensorValue `json:"te"`
		Hu sensorValue `json:"hu"`
		Il sensorValue `json:"il"`
		Mo sensorValue `json:"mo"`
	} `json:"newest_events"`
}

const url_devices = "https://api.nature.global/1/devices"

// TODO 作成中
func (s NatureRemoSdk) GetDevice() ([]Device, error) {
	var devices []Device

	body, err := s.request("GET", url_devices)
	if err != nil {
		return devices, err
	}

	err = json.Unmarshal(body, &devices)
	if err != nil {
		return devices, err
	}

	return devices, nil
}
