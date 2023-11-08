package token

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type authClaim struct {
	UserId uint `json:"userId"`
	jwt.RegisteredClaims
}

func GenerateToken(userId uint) (string, error) {
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

func extractToken(bearerToken string) string {
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func VerifyToken(bearerToken string) error {
	tokenString := extractToken(bearerToken)
	_, err := jwt.ParseWithClaims(tokenString, &authClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	return err
}

func GetUserIdFromToken(bearerToken string) (uint, error) {
	tokenString := extractToken(bearerToken)
	token, err := jwt.ParseWithClaims(tokenString, &authClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*authClaim); ok && token.Valid {
		return claims.UserId, nil
	}

	return 0, fmt.Errorf("invalid token")
}
