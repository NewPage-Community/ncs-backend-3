package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Config struct {
	ExpireTime int64
	SecretKey  string
}

// NewTokenString create a new valid JWT
func (config Config) NewTokenString(payload map[string]interface{}) (string, error) {
	payload["exp"] = time.Now().Add(time.Duration(config.ExpireTime)).Unix()
	payload["iss"] = "NewPage"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(payload))
	return token.SignedString([]byte(config.SecretKey))
}

// GetToken get a token from JWT string
func (config Config) GetTokenPayload(tokenString string) (Payload, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	// Success to return token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return Payload(claims), nil
	}

	err = fmt.Errorf("invalid token: %v", token.Claims)
	return nil, err
}
