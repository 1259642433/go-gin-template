package utils

import (
	"go-gin-template/pkg/setting"
	"time"

	jwt "github.com/dgrijalva/jwt-go"


)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	ID int `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(id int,username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		id,
		username,
		password,
		jwt.StandardClaims {
			NotBefore: nowTime.Unix() - 60,
			ExpiresAt : expireTime.Unix(),
			Issuer : "go-gin-template",
			IssuedAt: nowTime.Unix(),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}