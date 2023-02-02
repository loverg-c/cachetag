package repository

import (
	"database/sql"
	"log"
	m "tags-finder/Domain/model"
	"tags-finder/Infrastructure/Database/config"
	"time"
)

func NewTag(t *m.Tag) {
	if t == nil {
		log.Fatal(t)
	}
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()

	query := `
INSERT INTO tags (description, score, created_at, updated_at)
VALUES ($1,$2,$3,$4) 
RETURNING id;
`
	err := config.Db().QueryRow(query, t.Description, t.Score, t.CreatedAt, t.UpdatedAt).Scan(&t.Id)

	if err != nil {
		log.Fatal(err)
	}
}

func FindTagById(id int) *m.Tag {
	var tag m.Tag

	row := config.Db().QueryRow("SELECT * FROM tags WHERE id = $1;", id)
	err := row.Scan(&tag.Id, &tag.Description, &tag.Score, &tag.CreatedAt, &tag.UpdatedAt)

	if err != nil {
		log.Fatal(err)
	}

	return &tag
}

func GetAllTag() *m.Tags {

	query := `
SELECT id,
       description,
       score, 
       created_at,
       updated_at
FROM tags
`
	rows, err := config.Db().Query(query)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	return parseTagsRows(rows)
}

func parseTagsRows(rows *sql.Rows) *m.Tags {
	var ts m.Tags

	for rows.Next() {
		var t m.Tag
		if err := rows.Scan(&t.Id,
			&t.Description,
			&t.Score,
			&t.CreatedAt,
			&t.UpdatedAt); err != nil {
			log.Fatal(err)
		}
		ts = append(ts, t)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return &ts
}

func UpdateTag(tag *m.Tag) {
	tag.UpdatedAt = time.Now()

	stmt, err := config.Db().
		Prepare("UPDATE tags SET description=$1, score=$2, updated_at=$3 WHERE id=$4;")

	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(tag.Description, tag.Score, tag.UpdatedAt, tag.Id)

	if err != nil {
		log.Fatal(err)
	}
}

func DeleteTagById(id int) error {
	stmt, err := config.Db().Prepare("DELETE FROM tag WHERE id=$1;")

	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(id)

	return err
}
