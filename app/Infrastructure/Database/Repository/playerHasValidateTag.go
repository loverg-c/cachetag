package repository

import (
	"database/sql"
	"log"
	m "tags-finder/Domain/model"
	"tags-finder/Infrastructure/Database/config"
	"time"
)

func NewPlayerHasValidateTag(p *m.PlayerHasValidateTag) {
	if p == nil {
		log.Fatal(p)
	}
	p.ValidatedAt = time.Now()

	row := config.Db().QueryRow(
		`
INSERT INTO player_has_validate_tag (player_id, tag_id, validated_at)
VALUES ($1,$2,$3);`,
		p.PlayerId, p.TagId, p.ValidatedAt)

	if row.Err() != nil {
		log.Println(row.Err())
	}
}

func parsePlayerHasTagsRows(rows *sql.Rows) *m.PlayerHasValidateTags {
	var phvts m.PlayerHasValidateTags

	for rows.Next() {
		var phtv m.PlayerHasValidateTag
		if err := rows.Scan(&phtv.PlayerId, &phtv.TagId,
			&phtv.ValidatedAt); err != nil {
			log.Fatal(err)
		}
		phvts = append(phvts, phtv)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return &phvts
}

func FindByPlayerId(id int) *m.PlayerHasValidateTags {
	rows, err := config.Db().Query("SELECT * FROM player_has_validate_tag WHERE player_id = $1;", id)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	return parsePlayerHasTagsRows(rows)
}

func FindByTagId(id int) *m.PlayerHasValidateTags {
	rows, err := config.Db().Query("SELECT * FROM player_has_validate_tag WHERE tag_id = $1;", id)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	return parsePlayerHasTagsRows(rows)
}

func FindByPlayerAndTagId(player_id, tag_id int) (*m.PlayerHasValidateTag, error) {
	var pht m.PlayerHasValidateTag

	row := config.Db().QueryRow("SELECT player_id, tag_id, validated_at FROM player_has_validate_tag WHERE tag_id = $1 AND player_id = $2;", tag_id, player_id)
	err := row.Scan(&pht.PlayerId, &pht.TagId, &pht.ValidatedAt)

	if err != nil {
		return nil, err
	}

	return &pht, nil
}

func DeleteByPlayer(id int) error {
	stmt, err := config.Db().Prepare("DELETE FROM player_has_validate_tag WHERE player_id=$1;")

	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(id)

	return err
}

func DeleteByTag(id int) error {
	stmt, err := config.Db().Prepare("DELETE FROM player_has_validate_tag WHERE tag_id=$1;")

	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(id)

	return err
}

func DeleteByPlayerAndTag(player_id, tag_id int) error {
	stmt, err := config.Db().Prepare("DELETE FROM player_has_validate_tag WHERE tag_id=$1 AND player_id=$2;")

	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(tag_id, player_id)

	return err
}

func GetAllPlayerHasValidateTag() *m.PlayerHasValidateTags {
	rows, err := config.Db().Query("SELECT * FROM player_has_validate_tag")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	return parsePlayerHasTagsRows(rows)
}
