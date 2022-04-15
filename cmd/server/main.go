package main

import (
	"challenge/pkg/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

//Configuring env variables
func configureViper() error {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

//testing bitly link shortener
func main() {
	if err := configureViper(); err != nil {
		log.Fatal(err.Error())
	}

	token, ok := viper.Get("BITLY_OAUTH_TOKEN").(string)
	if !ok {
		log.Fatalf("no such env variable")
	}

	var data = strings.NewReader(`{ "long_url": "https://dev.bitly.com" }`)
	req, err := http.NewRequest("POST", "https://api-ssl.bitly.com/v4/shorten", data)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	bitlyResponse := &utils.BitlyResponse{}
	json.Unmarshal(bodyText, bitlyResponse)

	fmt.Printf("%s\n", bitlyResponse.ShortLink)
}
