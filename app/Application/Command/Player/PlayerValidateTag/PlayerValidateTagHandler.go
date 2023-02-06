package Command

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
	"tags-finder/Domain/model"
	repository "tags-finder/Infrastructure/Database/Repository"
	"tags-finder/Infrastructure/Token/JWT"
)

func HandlePlayerValidateTag(playerCommand PlayerValidateTag) {
	//todo create id before command
	playerVTag := model.PlayerHasValidateTag{
		PlayerId: playerCommand.PlayerId,
		TagId:    playerCommand.TagId,
	}

	repository.NewPlayerHasValidateTag(&playerVTag)

	jwt := JWT.GenerateJWT()
	var bearer = "Bearer " + jwt

	data, err := json.Marshal(*repository.GetScorePerPlayers())
	if err != nil {
		log.Println(err)
	}

	form := url.Values{}
	form.Add("topic", "http://localhost/tableau-de-score")
	form.Add("data", string(data))

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
