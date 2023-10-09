package model

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// ref. https://qiita.com/TakahikoKawasaki/items/970548727761f9e02bcd

type AccessToken struct {
	Iss      string   `json:"iss"`
	Sub      string   `json:"sub"`
	Exp      int64    `json:"exp"`
	Iat      int64    `json:"iat"`
	Aud      string   `json:"aud"`
	Jti      string   `json:"jti"`
	Scope    []string `json:"scope"`
	ClientID string   `json:"client_id"`
}

func GenerateAccessToken(
	iss string,
	sub string,
	aud string,
	jti string,
	scope []string,
	clientID string,
) AccessToken {
	now := time.Now()

	return AccessToken{
		Iss:      iss,
		Sub:      sub,
		Exp:      now.Add(time.Hour).Unix(),
		Iat:      now.Unix(),
		Aud:      aud,
		Jti:      jti,
		Scope:    scope,
		ClientID: clientID,
	}
}

func (at AccessToken) JWT(
	sign string,
) string {
	claims := jwt.MapClaims{
		"sub":       at.Sub,
		"iss":       at.Iss,
		"aud":       at.Aud,
		"exp":       at.Exp,
		"iat":       at.Iat,
		"jti":       at.Jti,
		"scope":     at.Scope,
		"client_id": at.ClientID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	str, _ := token.SignedString([]byte(sign))

	return str
}

func ParseAccessToken(
	str string,
	sign string,
) (AccessToken, error) {
	token, err := jwt.Parse(str, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return []byte(sign), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		sc := claims["scope"].([]interface{})
		scope := make([]string, len(sc))
		for i, v := range sc {
			scope[i] = v.(string)
		}

		return AccessToken{
			Iss:      claims["iss"].(string),
			Sub:      claims["sub"].(string),
			ClientID: claims["client_id"].(string),
			Exp:      int64(claims["exp"].(float64)),
			Iat:      int64(claims["iat"].(float64)),
			Scope:    scope,
			Aud:      claims["aud"].(string),
			Jti:      claims["jti"].(string),
		}, nil
	} else {
		return AccessToken{}, err
	}
}

type key struct{}

func SetAccessTokenCtx(
	ctx context.Context,
	val AccessToken,
) context.Context {
	return context.WithValue(ctx, key{}, val)
}

func GetAccessTokenCtx(
	ctx context.Context,
) AccessToken {
	v := ctx.Value(key{})

	vv, ok := v.(AccessToken)
	if !ok {
		return AccessToken{}
	}

	return vv
}
