package token

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type authClaim struct {
	UserId int `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(userId int) (string, error) {
	tokenLifeSpanHour, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_EXPIRE"))
	if err != nil {
		return "", err
	}

	customClaims := authClaim{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(tokenLifeSpanHour) * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func VerifyToken(tokenString string) error {
	_, err := jwt.ParseWithClaims(tokenString, &authClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	return err
}

func GetUserIdFromToken(tokenString string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenString, &authClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*authClaim); ok && token.Valid {
		return claims.UserId, nil
	}

	return 0, fmt.Errorf("Invalid token")
}
