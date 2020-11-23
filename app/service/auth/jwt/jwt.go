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

type Token struct {
	UID  int64 `json:"uid"`
	Flag int32 `json:"flag"`
	jwt.StandardClaims
}

// NewTokenString create a new valid JWT
func (t *Token) NewTokenString(config Config) (string, error) {
	t.ExpiresAt = time.Now().Add(time.Duration(config.ExpireTime)).Unix()
	t.Issuer = "NewPage"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, t)
	return token.SignedString(config.SecretKey)
}

// GetToken get a token from JWT string
func GetToken(tokenString string, config Config) (*Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return config.SecretKey, nil
	})

	// Success to return token
	if claims, ok := token.Claims.(Token); ok && token.Valid {
		return &claims, nil
	}

	// Invalid token
	if err == nil {
		err = fmt.Errorf("invalid token: %v", token.Claims)
	}
	return nil, err
}
