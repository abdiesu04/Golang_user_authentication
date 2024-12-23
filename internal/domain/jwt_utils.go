package domain

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTUtils struct {
	SecretKey string
}

func NewJWTUtils(secretKey string) *JWTUtils {
	return &JWTUtils{SecretKey: secretKey}
}

func (j *JWTUtils) GenerateToken(userID int64) (Token, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return Token{}, err
	}

	refreshClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(j.SecretKey))
	if err != nil {
		return Token{}, err
	}

	return Token{
		AccessToken:  accessToken,
		RefreshToken: refreshTokenString,
	}, nil
}
