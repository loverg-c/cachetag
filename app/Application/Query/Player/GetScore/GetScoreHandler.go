package GetScore

import (
	"tags-finder/Domain/model"
	repository "tags-finder/Infrastructure/Database/Repository"
)

func HandleGetScore(scoreQuery GetScore) *model.Score {
	score, err := repository.FindScoreByUser(scoreQuery.PlayerId)

	if err != nil {
		return nil
	}

	return score
}
