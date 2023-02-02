package Command

import (
	"tags-finder/Domain/model"
	repository "tags-finder/Infrastructure/Database/Repository"
)

func HandlePlayerValidateTag(playerCommand PlayerValidateTag) {
	//todo create id before command
	playerVTag := model.PlayerHasValidateTag{
		PlayerId: playerCommand.PlayerId,
		TagId:    playerCommand.TagId,
	}

	//todo check already validated tag here

	repository.NewPlayerHasValidateTag(&playerVTag)
}
