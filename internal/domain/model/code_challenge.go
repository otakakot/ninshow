package model

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/google/uuid"
)

type CodeChallengeMethod string

const (
	CodeChallengeMethodS256 CodeChallengeMethod = "S256"
)

type CodeChallenge struct {
	Challenge string
	Method    CodeChallengeMethod
}

func (cc CodeChallenge) Verify(verifier string) error {
	vel := sha256.Sum256([]byte(verifier))

	if base64.RawURLEncoding.EncodeToString(vel[:]) != cc.Challenge {
		return fmt.Errorf("code challenge does not match")
	}

	return nil
}

func GenerateCodeVerifier() string {
	b := make([]byte, 32)

	if _, err := rand.Read(b); err != nil {
		return base64.RawURLEncoding.EncodeToString([]byte(uuid.NewString()))
	}

	return base64.RawURLEncoding.EncodeToString(b)
}

func GenerateCodeChallenge(verifier string) CodeChallenge {
	vel := sha256.Sum256([]byte(verifier))

	return CodeChallenge{
		Challenge: base64.RawURLEncoding.EncodeToString(vel[:]),
		Method:    CodeChallengeMethodS256,
	}
}
