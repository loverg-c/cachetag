package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "database"
	port     = 5432
	user     = "app"
	password = "!ChangeMe!"
	dbname   = "app"
)

var db *sql.DB

func DatabaseInit() {
	var err error

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}

	createPlayerTable()
	createTagTable()
	createPlayerHasValidatedTagTable()
	createScoreView()
}

func createPlayerTable() {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS players(id serial,username varchar(35), created_at timestamp default NULL, updated_at timestamp default NULL, constraint pk_player primary key(id))")

	if err != nil {
		log.Fatal(err)
	}
}

func createTagTable() {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS tags(id serial,description varchar(50), score integer, created_at timestamp default NULL, updated_at timestamp default NULL, constraint pk_tag primary key(id))")

	if err != nil {
		log.Fatal(err)
	}
}

func createPlayerHasValidatedTagTable() {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS player_has_validate_tag(player_id int, tag_id int, validated_at timestamp default NULL, constraint fk_p foreign key(player_id) REFERENCES players(id), constraint fk_t foreign key(tag_id) REFERENCES tags(id))")

	if err != nil {
		log.Fatal(err)
	}
}

func createScoreView() {
	_, err := db.Exec(`
CREATE OR REPLACE VIEW score(player_id, player_username, score) AS
SELECT p.id                      AS player_id,
       p.username                AS player_username,
       sum(COALESCE(t.score, 0)) AS score
FROM players p
        LEFT JOIN player_has_validate_tag phvt ON phvt.player_id = p.id
        LEFT JOIN tags t ON t.id = phvt.tag_id
GROUP BY p.id, p.username;
`)

	if err != nil {
		log.Fatal(err)
	}
}

// Getter for db var
func Db() *sql.DB {
	return db
}
