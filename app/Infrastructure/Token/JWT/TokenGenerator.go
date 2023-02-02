package JWT

import (
	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"mercure": map[string]interface{}{
			"publish": []string{"*"},
		},
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString("!ChangeThisMercureHubJWTSecretKey!")

	return tokenString
}
