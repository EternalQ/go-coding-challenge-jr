package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

func BitlyShortener(link string) (string, error) {
	token, ok := viper.Get("BITLY_OAUTH_TOKEN").(string)
	if !ok {
		return "", errors.New("can't get bitly Oauth token")
	}

	var data = strings.NewReader(fmt.Sprintf(`{ "long_url": "%v" }`, link))
	req, err := http.NewRequest("POST", "https://api-ssl.bitly.com/v4/shorten", data)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return "", errors.New("may be wrong url or bitly api error")
	}
	defer resp.Body.Close()

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	bitlyResponse := &BitlyResponse{}
	json.Unmarshal(bodyText, bitlyResponse)

	return bitlyResponse.ShortLink, nil
}
