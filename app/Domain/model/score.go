package model

type Score struct {
	PlayerId int    `json:"player_id"`
	Username string `json:"username"`
	Score    int    `json:"score"`
}

type Scores []Score
