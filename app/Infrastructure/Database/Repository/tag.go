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
INSERT INTO tags (description, secret, score, created_at, updated_at)
VALUES ($1,md5(random()::text), $2,$3,$4,) 
RETURNING id;
`
	err := config.Db().QueryRow(query, t.Description, t.Secret, t.Score, t.CreatedAt, t.UpdatedAt).Scan(&t.Id)

	if err != nil {
		log.Fatal(err)
	}
}

func FindTagById(id int) (*m.Tag, error) {
	var tag m.Tag

	row := config.Db().QueryRow("SELECT * FROM tags WHERE id = $1;", id)
	err := row.Scan(&tag.Id, &tag.Description, &tag.Secret, &tag.Score, &tag.CreatedAt, &tag.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &tag, nil
}

func GetAllTag() *m.Tags {

	query := `
SELECT id,
       description,
       secret,
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
			&t.Secret,
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
		Prepare("UPDATE tags SET description=$1, secret=$2, score=$3, updated_at=$4 WHERE id=$5;")

	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(tag.Description, tag.Secret, tag.Score, tag.UpdatedAt, tag.Id)

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
