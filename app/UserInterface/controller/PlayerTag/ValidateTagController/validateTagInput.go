package ValidateTagController

type ValidateTagInput struct {
	PlayerId int    `validate:"required" json:"player_id"`
	TagId    int    `validate:"required" json:"tag_id"`
	Secret   string `validate:"required" json:"secret"`
}
