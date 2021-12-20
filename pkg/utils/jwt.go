package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

//var jwtSecret = []byte(setting.JwtSecret)

type CustomClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username, password string, secret []byte) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := CustomClaims{
		username,
		password,
		jwt.StandardClaims{
			//Audience:  "",
			ExpiresAt: expireTime.Unix(),
			//Id:        "",
			//IssuedAt:  expireTime.Unix(),
			Issuer: "pokertime",
			//NotBefore: expireTime.Unix(),
			//Subject:   "",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(secret)
}

func ParseToken(token string, secret []byte) (*CustomClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*CustomClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
