package repository

import (
	"database/sql"
	"log"
	m "tags-finder/Domain/model"
	"tags-finder/Infrastructure/Database/config"
	"time"
)

func NewPlayer(p *m.Player) {
	if p == nil {
		log.Fatal(p)
	}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	err := config.Db().QueryRow("INSERT INTO players (username, created_at, updated_at) VALUES ($1,$2,$3) RETURNING id;", p.Username, p.CreatedAt, p.UpdatedAt).Scan(&p.Id)

	if err != nil {
		log.Fatal(err)
	}
}

func FindPlayerById(id int) *m.Player {
	var player m.Player

	row := config.Db().QueryRow("SELECT * FROM players WHERE id = $1;", id)
	err := row.Scan(&player.Id, &player.Username, &player.CreatedAt, &player.UpdatedAt)

	if err != nil {
		log.Fatal(err)
	}

	return &player
}

func GetAllPlayer() *m.Players {
	rows, err := config.Db().Query("SELECT * FROM players")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	return parsePlayersRows(rows)
}

func parsePlayersRows(rows *sql.Rows) *m.Players {
	var ps m.Players

	for rows.Next() {
		var p m.Player
		if err := rows.Scan(&p.Id, &p.Username, &p.CreatedAt,
			&p.UpdatedAt); err != nil {
			log.Fatal(err)
		}
		ps = append(ps, p)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return &ps
}

func UpdatePlayer(player *m.Player) {
	player.UpdatedAt = time.Now()

	stmt, err := config.Db().Prepare("UPDATE players SET username=$1, updated_at=$2 WHERE id=$3;")

	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(player.Username, player.UpdatedAt, player.Id)

	if err != nil {
		log.Fatal(err)
	}
}

func DeletePlayerById(id int) error {
	stmt, err := config.Db().Prepare("DELETE FROM player WHERE id=$1;")

	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(id)

	return err
}
