package tools

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtKeyString string

var (
	JwtCtxKey = jwtKeyString("jwt-auth-key")
)

type JwtCustomClaim struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

var jwtSecret = []byte(getJwtSecret())

func getJwtSecret() string {
	secret := os.Getenv("JWT_SECRET")

	return secret
}

func JwtGenerate(userID int) string {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &JwtCustomClaim{
		ID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	token, err := t.SignedString(jwtSecret)
	if err != nil {
		panic(err)
	}

	return token
}

func JwtValidate(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &JwtCustomClaim{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there's a problem with the signing method")
		}
		return jwtSecret, nil
	})
}

func AuthCtx(ctx context.Context) *JwtCustomClaim {
	raw, _ := ctx.Value(JwtCtxKey).(*JwtCustomClaim)
	return raw
}
