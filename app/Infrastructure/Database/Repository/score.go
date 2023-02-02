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
