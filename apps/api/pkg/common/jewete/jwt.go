package jewete

import (
	"fmt"

	jwt "github.com/golang-jwt/jwt/v5"
)

type JWTResult struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

const JWT_KEY = "inisaltjwt" // TODO: Update This move into env variable

func JWTEncrypt(data any) (*string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"data": data,
		})
	s, err := t.SignedString([]byte(JWT_KEY))
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func JWTDecrypt(token string) (any, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JWT_KEY), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("failed to get claims")
	}
	return claims["data"], nil
}
