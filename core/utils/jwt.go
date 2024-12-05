package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

var secretKey = []byte("secret")

func GenerateJWT(userID int64) (string, string, error) {
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

func VerifyRefreshToken(userToken string) (int64, error) {
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(
		userToken,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		},
	)
	if err != nil || !token.Valid {
		log.Println("error occurred in verify refresh token :", err)
		return 0, err
	}

	// Extract user_id from the claims
	userIDFloat, ok := claims["user_id"].(float64) // Claims are unmarshalled as float64
	if !ok {
		log.Println("error occurred in verify refresh token : user_id not found in token claims")
		return 0, errors.New("user_id not found in token claims")
	}

	userID := int64(userIDFloat)
	return userID, nil
}

func VerifyToken(userToken string) (int64, error) {
	// Parse the token with claims
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(
		userToken,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		},
	)
	if err != nil || !token.Valid {
		log.Println("error occurred in verify token:", err)
		return 0, err
	}

	// Check expiration time
	exp, ok := claims["exp"].(float64)
	if !ok {
		log.Println("error occurred in verify token: exp claim missing or invalid")
		return 0, errors.New("invalid token: expiration claim missing")
	}
	if time.Unix(int64(exp), 0).Before(time.Now()) {
		log.Println("error occurred in verify token: token has expired")
		return 0, errors.New("token has expired")
	}

	// Extract user_id from the claims
	userIDFloat, ok := claims["user_id"].(float64) // Claims are unmarshalled as float64
	if !ok {
		log.Println("error occurred in verify token: user_id not found in token claims")
		return 0, errors.New("user_id not found in token claims")
	}

	userID := int64(userIDFloat)
	return userID, nil
}
