package timer

import (
	"fmt"
	"net/http"
)

const (
	timerApi string = "https://timercheck.io/"
)

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
func CheckTimerAPI(name string) (bool, error) {
	url := fmt.Sprintf("%s%s", timerApi, name)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil || res.StatusCode != 200 {
		return false, err
	}

	return true, nil
}
