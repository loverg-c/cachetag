package controller

type CreatePlayerInput struct {
	Username string `validate:"required"`
}
