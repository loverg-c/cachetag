package repository

import (
	"log"
	m "tags-finder/Domain/model"
	"tags-finder/Infrastructure/Database/config"
)

func GetScorePerPlayers() *m.Scores {
	query := `
SELECT player_id, player_username, score
FROM score
`

	rows, err := config.Db().Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var scoreList m.Scores

	for rows.Next() {
		var score m.Score
		if err := rows.Scan(&score.PlayerId, &score.Username, &score.Score); err != nil {
			log.Fatal(err)
		}
		scoreList = append(scoreList, score)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return &scoreList
}

func FindScoreByUser(playerId int) (*m.Score, error) {
	query := `
SELECT player_id, player_username, score
FROM score
WHERE player_id = $1;
`
	var score m.Score

	row := config.Db().QueryRow(query, playerId)
	err := row.Scan(&score.PlayerId, &score.Username, &score.Score)

	if err != nil {
		return nil, err
	}

	return &score, nil
}
