package Query

import (
	"tags-finder/Domain/model"
	repository "tags-finder/Infrastructure/Database/Repository"
)

func HandleGetPlayer(playerQuery GetPlayer) *model.Player {
	player, err := repository.FindPlayerByUsername(playerQuery.Username)

	if err != nil {
		return nil
	}

	return player
}
