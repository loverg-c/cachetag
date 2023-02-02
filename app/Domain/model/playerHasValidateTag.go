package model

import (
	"time"
)

type PlayerHasValidateTag struct {
	PlayerId    int       `json:"player_id"`
	TagId       int       `json:"tag_id"`
	ValidatedAt time.Time `json:"validated_at"`
}

type PlayerHasValidateTags []PlayerHasValidateTag
