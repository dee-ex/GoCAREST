package tokens

import (
	"fmt"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
)

// ValidateToken validates token string and extract a jwt token correspond with
func ValidateToken(tokenStr string) (*jwt.Token, error) {
	token, err := VerifyToken(tokenStr)
	if err != nil {
		return nil, fmt.Errorf("Token is unverified")
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return nil, fmt.Errorf("Token is invalid")
	}

	return token, nil
}

// VerifyToken verifies jwt token after validation stage
func VerifyToken(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}
