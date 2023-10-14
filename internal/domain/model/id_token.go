package model

import (
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// ref. https://qiita.com/TakahikoKawasaki/items/8f0e422c7edd2d220e06

// ref. https://openid.net/specs/openid-connect-core-1_0.html#StandardClaims
type IDToken struct {
	Iss     string  `json:"iss"`
	Sub     string  `json:"sub"`
	Aud     string  `json:"aud"`
	Nonce   string  `json:"nonce"`
	Exp     int64   `json:"exp"`
	Iat     int64   `json:"iat"`
	Profile *string `json:"profile"`
	Email   *string `json:"email"`
}

func GenerateIDToken(
	iss string,
	sub string,
	aud string,
	nonce string,
	profile *string,
	email *string,
) IDToken {
	now := time.Now()

	idt := IDToken{
		Iss:   iss,
		Sub:   sub,
		Aud:   aud,
		Nonce: nonce,
		Exp:   now.Add(time.Hour).Unix(),
		Iat:   now.Unix(),
	}

	if profile != nil {
		idt.Profile = profile
	}

	if email != nil {
		idt.Email = email
	}

	return idt
}

func (it IDToken) JWT(
	sign string,
) string {
	claims := jwt.MapClaims{
		"iss":   it.Iss,
		"sub":   it.Sub,
		"aud":   it.Aud,
		"exp":   it.Exp,
		"iat":   it.Iat,
		"nonce": it.Nonce,
	}

	if it.Profile != nil {
		claims["profile"] = it.Profile
	}

	if it.Email != nil {
		claims["email"] = it.Email
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	str, _ := token.SignedString([]byte(sign))

	return str
}

func (it IDToken) RSA256(
	key *rsa.PrivateKey,
) string {
	claims := jwt.MapClaims{
		"iss":   it.Iss,
		"sub":   it.Sub,
		"aud":   it.Aud,
		"exp":   it.Exp,
		"iat":   it.Iat,
		"nonce": it.Nonce,
	}

	if it.Profile != nil {
		claims["profile"] = it.Profile
	}

	if it.Email != nil {
		claims["email"] = it.Email
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	token.Header["kid"] = "12345678"

	str, _ := token.SignedString(key)

	return str
}

func ParseIDToken(
	str string,
	sign string,
) (IDToken, error) {
	token, err := jwt.Parse(str, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return []byte(sign), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		idt := IDToken{
			Iss:   claims["iss"].(string),
			Sub:   claims["sub"].(string),
			Exp:   claims["exp"].(int64),
			Iat:   claims["iat"].(int64),
			Aud:   claims["aud"].(string),
			Nonce: claims["nonce"].(string),
		}

		if profile, ok := claims["profile"].(string); ok {
			idt.Profile = &profile
		}

		if email, ok := claims["email"].(string); ok {
			idt.Email = &email
		}

		return idt, nil
	} else {
		return IDToken{}, err
	}
}