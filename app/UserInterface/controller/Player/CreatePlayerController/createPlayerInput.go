package CreatePlayerController

type CreatePlayerInput struct {
	Username string `validate:"required"`
}
