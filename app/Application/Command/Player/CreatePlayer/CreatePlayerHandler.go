package Command

import (
	"tags-finder/Domain/model"
	repository "tags-finder/Infrastructure/Database/Repository"
)

func HandleCreatePlayer(playerCommand CreatePlayer) model.Player {
	//todo create id before command
	player := model.Player{
		Username: playerCommand.Username,
	}

	repository.NewPlayer(&player)

	return player
}
