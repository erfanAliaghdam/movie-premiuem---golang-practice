package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var secretKey = []byte("secret")

func GenerateJWT(userID string) (string, string, error) {
	AccessClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"iat":     time.Now().Unix(),
	}

	AccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, AccessClaims)

	signedAccessToken, signAccessTokenError := AccessToken.SignedString(secretKey)
	if signAccessTokenError != nil {
		return "", "", signAccessTokenError
	}

	RefreshClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
		"iat":     time.Now().Unix(),
	}

	RefreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, RefreshClaims)

	signedRefreshToken, signRefreshTokenError := RefreshToken.SignedString(secretKey)
	if signRefreshTokenError != nil {
		return "", "", signRefreshTokenError
	}

	return signedAccessToken, signedRefreshToken, nil
}

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func VerifyToken(userToken string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(
		userToken,
		*claims,
		func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		},
	)
	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
