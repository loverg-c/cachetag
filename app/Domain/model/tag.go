package model

import (
	"time"
)

type Tag struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Secret      string    `json:"-"`
	Score       int       `json:"score"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Tags []Tag
