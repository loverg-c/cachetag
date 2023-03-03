package Mercure

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
	"tags-finder/Infrastructure/Token/JWT"
)

func SendNotificationWithMercure(topic string, data interface{}) {
	jwt := JWT.GenerateJWT()
	var bearer = "Bearer " + jwt

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}

	form := url.Values{}
	form.Add("topic", topic)
	form.Add("data", string(jsonData))

	req, err := http.NewRequest(
		http.MethodPost,
		"http://mercure/.well-known/mercure",
		strings.NewReader(form.Encode()))

	if err != nil {
		log.Println(err)
	}

	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
}
