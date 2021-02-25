package tokens

import (
	"errors"

	jwt "github.com/dgrijalva/jwt-go"
)

// ExtractClaims extracts claims from token
func ExtractClaims(token *jwt.Token) (jwt.MapClaims, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("Cannot extract claims from token")
	}

	return claims, nil
}
