package common

import (
	"context"
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

var CtxKeyJwtUserId = "userId"

func GetJwtToken(secret string, nowDate int64, expire int64, userId string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = nowDate + expire
	claims["startFrom"] = nowDate
	claims[CtxKeyJwtUserId] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secret))
}

func GetUidFromCtx(ctx context.Context) string {
	// var uid string
	if uid, ok := ctx.Value(CtxKeyJwtUserId).(string); ok {
		return uid
	} else {
		return ""
	}
}

func ParseTokenStr(tokenStr string, secret string) (*jwt.MapClaims, error) {
	// exp := l.ctx.Value("exp")
	// startFrom := l.ctx.Value("startFrom")
	// userId := l.ctx.Value("userId")

	token, err := jwt.ParseWithClaims(tokenStr, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("parse invalid token")
	}

	res, ok := token.Claims.(*jwt.MapClaims)
	if !ok {
		return nil, errors.New("jwt.MapClaims invalid")
	}

	return res, nil
}
