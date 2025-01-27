package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	timerApi string = "https://timercheck.io/"
)

type timerApiResponse struct {
	Error   string  `json:"errorMessage"`
	Name    string  `json:"timer"`
	Seconds float64 `json:"seconds_remaining"`
}

// Create or update API timer
func StartTimerAPI(name string, seconds int64) error {
	url := fmt.Sprintf("%s%s/%v", timerApi, name, seconds)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return err
	}

	return nil
}

// Check existing timer
func CheckTimerAPI(name string) (*timerApiResponse, error) {
	url := fmt.Sprintf("%s%s", timerApi, name)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	timerRes := &timerApiResponse{}
	json.NewDecoder(res.Body).Decode(timerRes)

	return timerRes, nil
}
