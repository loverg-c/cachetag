package VerifyTagController

type VerifyTagInput struct {
	Secret string `validate:"required"`
}
