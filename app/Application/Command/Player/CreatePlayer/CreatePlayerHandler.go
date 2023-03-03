package Command

import (
	"tags-finder/Domain/model"
	repository "tags-finder/Infrastructure/Database/Repository"
	"tags-finder/Infrastructure/Notification/Mercure"
)

func HandleCreatePlayer(playerCommand CreatePlayer) model.Player {
	//todo create id before command
	player := model.Player{
		Username: playerCommand.Username,
	}

	repository.NewPlayer(&player)

	Mercure.SendNotificationWithMercure(
		"http://localhost/tableau-de-score",
		*repository.GetScorePerPlayers())

	return player
}
