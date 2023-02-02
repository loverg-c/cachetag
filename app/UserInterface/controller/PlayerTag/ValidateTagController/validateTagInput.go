package ValidateTagController

type ValidateTagInput struct {
	PlayerId int `validate:"required"`
	TagId    int `validate:"required"`
}
