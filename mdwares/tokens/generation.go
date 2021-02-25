package tokens

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CreateAuthToken creates token string from data
// lifetime scale in minute
func CreateAuthToken(data map[string]interface{}, lifetime int) (string, error) {
	// Claim the token
	token := jwt.New(jwt.SigningMethodHS256)
	tokenClaims := token.Claims.(jwt.MapClaims)
	tokenClaims["authorized"] = true
	tokenClaims["exp"] = time.Now().Add(time.Duration(lifetime) * time.Minute).Unix()

	// Insert data to claims for token
	for k, v := range data {
		tokenClaims[k] = v
	}

	// Create the token
	authToken, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return authToken, nil
}
