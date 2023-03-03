package Command

import (
	"tags-finder/Domain/model"
	repository "tags-finder/Infrastructure/Database/Repository"
	"tags-finder/Infrastructure/Notification/Mercure"
)

func HandlePlayerValidateTag(playerCommand PlayerValidateTag) {
	//todo create id before command
	playerVTag := model.PlayerHasValidateTag{
		PlayerId: playerCommand.PlayerId,
		TagId:    playerCommand.TagId,
	}

	repository.NewPlayerHasValidateTag(&playerVTag)

	Mercure.SendNotificationWithMercure(
		"http://localhost/tableau-de-score",
		*repository.GetScorePerPlayers())
}
