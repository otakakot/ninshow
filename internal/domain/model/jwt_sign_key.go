package model

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/binary"
	"fmt"

	"github.com/google/uuid"
)

type JWTSignKey struct {
	ID  string
	Key *rsa.PrivateKey
}

func GenerateJWTSignKey() (*JWTSignKey, error) {
	reader := rand.Reader

	bitSize := 2048

	key, err := rsa.GenerateKey(reader, bitSize)
	if err != nil {
		return nil, fmt.Errorf("failed to generate key: %v", err)
	}

	return &JWTSignKey{
		ID:  uuid.New().String(),
		Key: key,
	}, nil
}

func (jsk JWTSignKey) Cert() Cert {
	data := make([]byte, 8)

	binary.BigEndian.PutUint64(data, uint64(jsk.Key.PublicKey.E))

	i := 0
	for ; i < len(data); i++ {
		if data[i] != 0x0 {
			break
		}
	}

	e := base64.RawURLEncoding.EncodeToString(data[i:])

	return Cert{
		Kid: jsk.ID,
		Kty: "RSA",
		Use: "sig",
		Alg: "RS256",
		N:   base64.RawURLEncoding.EncodeToString(jsk.Key.PublicKey.N.Bytes()),
		E:   e,
	}
}
