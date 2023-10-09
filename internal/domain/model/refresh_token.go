package model

import (
	"encoding/base64"

	"github.com/google/uuid"
)

type RefreshToken string

func GenerateRefreshToken() RefreshToken {
	token := uuid.NewString()

	return RefreshToken(token)
}

func (rt RefreshToken) Base64() string {
	return base64.StdEncoding.EncodeToString([]byte(rt))
}

func ParseRefreshToken(str string) (RefreshToken, error) {
	token, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}

	return RefreshToken(token), nil
}
